lexer grammar GrobbitLexer;

TtLParen:           '(';
TtRParen:           ')';
TtLBrace:           '{';
TtRBrace:           '}';
TtComma:            ',';
TtPeriod:           '.';
TtSemicolon:        ';';

TtOpAdd:            '+';
TtOpSub:            '-';
TtOpMul:            '*';
TtOpDiv:            '/';
TtOpMod:            '%';
TtOpAssign:         '=';
TtOpAnd:            '&&';
TtOpOr:             '||';
TtOpNot:            '!';
TtOpEq:             '==';
TtOpLt:             '<';
TtOpGt:             '>';
TtOpNe:             '!=';
TtOpLe:             '<=';
TtOpGe:             '>=';
TtOpDefine:         ':=';

TtKwBreak:      'break';
TtKwCase:       'case';
TtKwChan:       'chan';
TtKwConst:      'const';
TtKwContinue:   'continue';
TtKwDefault:    'default';
TtKwDefer:      'defer';
TtKwElse:       'else';
TtKwFallthrough: 'fallthrough';
TtKwFor:        'for';
TtKwFunc:       'func';
TtKwGo:         'go';
TtKwGoto:       'goto';
TtKwIf:         'if';
TtKwImport:     'import';
TtKwInterface:  'interface';
TtKwMap:        'map';
TtKwPackage:    'package';
TtKwRange:      'range';
TtKwReturn:     'return';
TtKwSelect:     'select';
TtKwStruct:     'struct';
TtKwSwitch:     'switch';
TtKwType:       'type';
TtKwVar:        'var';

TtBool:         'false' | 'true';

TtIdentifier:   [a-zA-Z_][a-zA-Z_0-9]*; // NOTE: unicode?

TtFloat
    : DEC_DIGITS '.' DEC_DIGITS? EXPONENT?
    | DEC_DIGITS EXPONENT
    | '.' DEC_DIGITS EXPONENT?
    ;

TtInt:          [1-9][0-9]* | [0];

TtString:       '"' (EscapeInStr | ~[\r\n])* '"';
TtChar:         '\'' (EscapeInChar | ~[\r\n]) '\'';

// Fragment rules to build the float syntax (No separators)
fragment DEC_DIGITS
    : [0-9]+
    ;

fragment EXPONENT
    : [eE] [+-]? DEC_DIGITS
    ;

fragment EscapeInStr: '\\' ( '"' | '\\' | 'n' | 't');
fragment EscapeInChar: '\\' ( '\'' | '\\' |  'n' | 't');

// Whitespaces and comments (to be skipped)
WS : [ \t\r\n]+ -> skip;
BLOCK_COMMENT : '/*' .*? '*/' -> skip;
LINE_COMMENT : '//' ~[\r\n]* -> skip;

TtUnknown : . ;
