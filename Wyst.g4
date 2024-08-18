grammar Wyst;

WS: [ \t\r\n]+ -> skip;
NUMBER: [1-9][0-9]* | '0';
HEX: '0x' [a-fA-F0-9]*;
IDENTIFIER: [a-zA-Z_] ([a-zA-Z0-9] | '::')*;
MATH: '+' | '-' | '*' | '/';
fragment ESC: '\\' ['"\\] ;
STRING: '"' (ESC | ~["\\])* '"';
round_def: '(' (var_def (',' var_def)* ','?)? ')';
func_def: IDENTIFIER IDENTIFIER round_def code_block;
var_def: IDENTIFIER IDENTIFIER;
var_def_set: var_def '=' expr;
asm: '%' IDENTIFIER expr (',' expr)* ';';

expr:
    (   NUMBER |
        HEX |
        IDENTIFIER |
        STRING
    ) (MATH expr)*;

code_block: '{' ((
    expr |
    var_def |
    var_def_set
) ';')* '}';
top: (asm|func_def)*;