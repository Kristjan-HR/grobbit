package llvm

import (
	"fmt"
	"os"
	"slices"
	"strings"
)

///////////////////////////////////////////////////////////////////////////////////////////////////////////////////

// Block represents a basic block. It always has an entry label and a terminating instructions.
type Block struct {
	owner    *Function
	lab      Label
	comments []string
	instrs   []Instruction
	isClosed bool
}

// GetLabel returns the (entry) label of the block.
func (b *Block) GetLabel() Label          { return b.lab }
func (b *Block) SetLabel(lab Label)       { b.lab = lab }
func (b *Block) IsClosed() bool           { return b.isClosed }
func (b *Block) GetOwner() *Function      { return b.owner }
func (b *Block) SetOwner(owner *Function) { b.owner = owner }

// String returns a string representation of the block.
func (b *Block) String() string {
	var sb strings.Builder
	sb.WriteString(strings.Trim(b.lab.String(), "%"))
	sb.WriteString(":\n")
	for _, comment := range b.comments {
		sb.WriteString("  ; ")
		sb.WriteString(comment)
		sb.WriteString("\n")
	}
	for _, instr := range b.instrs {
		sb.WriteString("  ")
		sb.WriteString(instr.String())
		sb.WriteString("\n")
	}
	return sb.String()
}

// Prepend adds instruction(s) to the beginning of the block.
// It returns the register of last inserted instruction.
func (b *Block) Prepend(instrs ...Instruction) *Register {
	if b.IsClosed() {
		fmt.Println("Ops, trying to prepend to a closed block.")
		return nil
	}
	for _, instr := range instrs {
		if instr.IsTerminal() {
			fmt.Println("Ops, trying to insert a terminal instruction using prepend.")
			return nil
		}
	}
	b.instrs = slices.Insert(b.instrs, 0, instrs...)
	return instrs[len(instrs)-1].Register()
}

// Append adds instruction(s) to the end of the block.
// It returns the register of last inserted instruction.
func (b *Block) Append(instrs ...Instruction) *Register {
	if b.IsClosed() {
		fmt.Println("Ops, trying to append to a closed block.")
		return nil
	}
	for _, instr := range instrs {
		if instr.IsTerminal() {
			fmt.Println("Ops, trying to insert a terminal instruction using append.")
			return nil
		}
	}
	b.instrs = append(b.instrs, instrs...)
	return instrs[len(instrs)-1].Register()
}

// Close closes the block by setting its terminal instruction (no more modifications possible).
// It returns the register of inserted instruction.
func (b *Block) Close(instr Instruction) *Register {
	if !instr.IsTerminal() {
		fmt.Println("Ops, trying to insert a non-terminal instruction as terminal.")
		return nil
	}
	b.instrs = append(b.instrs, instr)
	b.isClosed = true
	return b.instrs[len(b.instrs)-1].Register()
}

// SplitBlock splits the current block into two blocks (with same owner, if any).
// If the closeFirst parameter is true, then the former block is closed with a branch instruction to the second.
// The function returns the latter block.
func (b *Block) SplitBlock(closeFirst bool) *Block {
	if b.IsClosed() {
		fmt.Println("Ops, trying to split an already closed block.")
		return nil
	}
	label := NewLabel()
	block := NewBlock(label)
	if closeFirst {
		b.Close(NewBranchInstr(label))
	}
	if b.owner != nil {
		b.owner.Append(block)
	}
	return block
}

// NewBlock creates a new block with a given label.
func NewBlock(label Label, comments ...string) *Block {
	return &Block{lab: label, comments: comments}
}

///////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

// Declaration is an interface for declaring global const and variables, and functions (incl. externals)
type Declaration interface {
	_Declaration_()
	String() string
}

type Global struct {
	Name    string
	Type    IRType
	Value   string
	IsConst bool
}

func (d *Global) _Declaration_() {}

func (d *Global) String() string {
	attribute := ""
	if d.IsConst {
		attribute = "constant"
	} else {
		attribute = "global"
	}
	return fmt.Sprintf("@%s = %s %s %s", d.Name, attribute, d.Type.String(), d.Value)
}

// NewGlobal returns a new Global (no owner).
func NewGlobal(name string, typ IRType, value string, isConst bool) *Global {
	return &Global{Name: name, Type: typ, Value: value, IsConst: isConst}
}

type Function struct {
	Prg    *Module
	Name   string
	Params []*Register
	Ret    IRType
	Blocks []*Block
}

func (d *Function) _Declaration_() {}

// String returns a string representation of a function declaration.
func (d *Function) String() string {
	var sb strings.Builder
	sb.WriteString("define ")
	sb.WriteString(d.Ret.String())
	sb.WriteString(" ")
	sb.WriteString(d.Name)
	sb.WriteString("(")
	for i, b := range d.Params {
		if i > 0 {
			sb.WriteString(", ")
		}
		sb.WriteString(b.String())
	}
	sb.WriteString(") {\n")
	for _, b := range d.Blocks {
		sb.WriteString(b.String())
	}
	sb.WriteString("}")
	return sb.String()
}

// LocateBlock returns the block with a given label, otherwise nil if not found.
func (d *Function) LocateBlock(lab Label) *Block {
	for _, b := range d.Blocks {
		if b.GetLabel() == lab {
			return b
		}
	}
	return nil
}

// Append adds a block to the end of function's block list (setting the function its owner).
// Returns the block added.
func (d *Function) Append(block *Block) *Block {
	if d.LocateBlock(block.GetLabel()) != nil {
		fmt.Println("Ops, trying to add a function block with the same label as existing block.")
		return nil
	}
	block.SetOwner(d)
	d.Blocks = append(d.Blocks, block)
	return block
}

// Prepend adds a block to the beginning of the function's block list (setting the function its owner).
// Returns the block added.
func (d *Function) Prepend(block *Block) *Block {
	if d.LocateBlock(block.GetLabel()) != nil {
		fmt.Println("Ops, trying to add a function block with the same label as existing block.")
		return nil
	}
	block.SetOwner(d)
	d.Blocks = slices.Insert(d.Blocks, 0, block)
	return block
}

// Relabel does rename all registers and labels of the function to be compliant with what LLVM-IR rules expect.
func (d *Function) Relabel() {
	var labelMap = make(map[uint]uint)
	var renamed = make(map[*Register]bool)

	var rename = func(r *Register) {
		if r == nil {
			return
		}
		if renamed[r] {
			return
		}
		r.Num = labelMap[r.Num]
		renamed[r] = true
	}

	var renameOperand = func(o Operand) {
		if r, ok := o.(*Register); ok {
			rename(r)
		}
	}

	var no uint = 0
	// map register and label values
	for _, param := range d.Params {
		labelMap[param.Num] = no
		no++
	}
	for _, block := range d.Blocks {
		labelMap[uint(block.GetLabel())] = no
		no++
		for _, instr := range block.instrs {
			reg := instr.Register()
			if reg != nil {
				labelMap[reg.Num] = no
				no++
			}
		}
	}
	// relabel
	for _, param := range d.Params {
		rename(param)
	}
	for _, block := range d.Blocks {
		block.lab = Label(labelMap[uint(block.GetLabel())])
		for _, instr := range block.instrs {
			reg := instr.Register()
			if reg != nil {
				rename(reg)
			}
			o := instr.GetO1()
			if o != nil {
				renameOperand(o)
			}
			o = instr.GetO2()
			if o != nil {
				renameOperand(o)
			}
			switch ist := instr.(type) {
			case *CallInstr:
				for i, _ := range ist.Args {
					renameOperand(ist.Args[i])
				}
			case *BranchInstr:
				ist.Lab = Label(labelMap[uint(ist.Lab)])
			case *BranchCondInstr:
				ist.TrueLab = Label(labelMap[uint(ist.TrueLab)])
				ist.FalseLab = Label(labelMap[uint(ist.FalseLab)])
			case *SelectInstr:
				renameOperand(ist.Cond)
			case *PhiInstr:
				for i, _ := range ist.options {
					renameOperand(ist.options[i].Op)
					ist.options[i].Label = Label(labelMap[uint(ist.options[i].Label)])
				}
			}
		}
	}
}

// SetName sets the name of the function.
func (d *Function) SetName(name string) {
	d.Name = name
}

// SetParams sets the parameters of the function.
func (d *Function) SetParams(params []*Register) {
	d.Params = params
}

// SetType sets the return type of the function.
func (d *Function) SetType(ret IRType) {
	d.Ret = ret
}

// NewFunction returns a new function declaration (no owner).
func NewFunction(name string, params []*Register, ret IRType) *Function {
	fn := &Function{Name: "@" + name, Params: params, Ret: ret, Blocks: nil}
	return fn
}

///////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

// Module is the top compiler unit (file) of the LLVM code.
type Module struct {
	Header []string
	Decls  []Declaration
}

// NewModule returns a new module.
func NewModule() *Module {
	return &Module{}
}

// Append appends a declaration to the module.
func (m *Module) Append(d Declaration) {
	m.Decls = append(m.Decls, d)
}

// AppendHeader appends a new string to the header.
func (m *Module) AppendHeader(s string) {
	m.Header = append(m.Header, s)
}

// String returns a string representation of a module.
func (m *Module) String() string {
	var sb strings.Builder
	for _, h := range m.Header {
		sb.WriteString(h)
		sb.WriteString("\n")
	}
	for _, d := range m.Decls {
		sb.WriteString(d.String())
		sb.WriteString("\n\n")
	}
	return sb.String()
}

// WriteToFile writes a string representation of the module to a file.
func (m *Module) WriteToFile(filename string) error {
	return os.WriteFile(filename, []byte(m.String()), 0644)
}

// Relabel does rename all registers and labels of the module's functions to be compliant with LLVM-IR.
func (m *Module) Relabel() {
	for _, d := range m.Decls {
		if funcDecl, ok := d.(*Function); ok {
			funcDecl.Relabel()
		}
	}
}
