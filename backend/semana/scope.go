// The code is adapted (simplified) from go/types2 (in Go compiler implementation).

package semana

import (
	"fmt"
	"io"
	"slices"
	"strings"
)

type Scope struct {
	parent   *Scope
	children []*Scope
	elems    map[string]Object
	comment  string // for debugging only
	isFunc   bool
}

func NewScope(parent *Scope, comment string) *Scope {
	s := &Scope{parent: parent, children: nil, elems: nil, comment: comment, isFunc: false}
	if parent != nil {
		parent.children = append(parent.children, s)
	}
	return s
}

func (s *Scope) Parent() *Scope { return s.parent }

func (s *Scope) Len() int { return len(s.elems) }

func (s *Scope) Names() []string {
	names := make([]string, len(s.elems))
	i := 0
	for name := range s.elems {
		names[i] = name
		i++
	}
	slices.Sort(names)
	return names
}

func (s *Scope) NumChildren() int { return len(s.children) }

func (s *Scope) Child(i int) *Scope { return s.children[i] }

func (s *Scope) Lookup(name string) Object {
	return s.elems[name]
}

func (s *Scope) Comment() string { return s.comment }

// Insert attempts to insert an object obj into scope s.
// If s already contains an alternative object alt with
// the same name, Insert leaves s unchanged and returns alt.
// Otherwise, it inserts obj, sets the object's parent scope
// if not already set, and returns nil.
func (s *Scope) Insert(obj Object) Object {
	name := obj.Name()
	if alt := s.Lookup(name); alt != nil {
		return alt
	}
	s.insert(name, obj)
	return nil
}

func (s *Scope) IsFunc() bool { return s.isFunc }

func (s *Scope) insert(name string, obj Object) {
	if s.elems == nil {
		s.elems = make(map[string]Object)
	}
	s.elems[name] = obj
}

func (s *Scope) WriteTo(w io.Writer, n int, recurse bool) {
	const ind = ".  "
	indn := strings.Repeat(ind, n)

	fmt.Fprintf(w, "%s%s scope %p {\n", indn, s.comment, s)

	indn1 := indn + ind
	for _, name := range s.Names() {
		fmt.Fprintf(w, "%s%s\n", indn1, s.Lookup(name))
	}

	if recurse {
		for _, s := range s.children {
			s.WriteTo(w, n+1, recurse)
		}
	}

	fmt.Fprintf(w, "%s}\n", indn)
}

func (s *Scope) String() string {
	var buf strings.Builder
	s.WriteTo(&buf, 0, true) // YB recurse
	return buf.String()
}
