grammar Wyst;

WS: [ \t\r\n]+ -> skip;
NUMBER: [1-9][0-9]* | '0';
HEX: '0x' [a-fA-F0-9]*;
IDENTIFIER: [a-zA-Z_] ([a-zA-Z0-9_] | '::')*;
MATH: '+' | '-' | '*' | '/';
fragment ESC: '\\' ['"\\] ;
STRING: '"' (ESC | ~["\\])* '"';
round_def: '(' (var_def (',' var_def)* ','?)? ')';
enum_curly: '{' (var_def ';')* '}';
round_call: '(' (expr (',' expr)* ','?)? ')';
fn_call: IDENTIFIER round_call;
func_def: IDENTIFIER IDENTIFIER round_def code_block;
var_def: IDENTIFIER IDENTIFIER;
var_def_set: var_def '=' expr;
call_tree: (fn_call|IDENTIFIER) ('.' (fn_call|IDENTIFIER))*;
struct_def: 'struct' IDENTIFIER enum_curly;
power_keyword: 'map';
power_keyword_call: '#' power_keyword expr (',' expr)* ';';
namespace: 'namespace' IDENTIFIER '{' (((var_def|var_def_set) ';')|func_def)* '}';

expr:
    (
        call_tree |
        NUMBER |
        HEX |
        IDENTIFIER |
        STRING
    ) (MATH expr)*;

code_block: '{' ((
    expr |
    var_def |
    var_def_set
) ';')* '}';
top: (
    power_keyword_call |
    func_def |
    struct_def |
    namespace
)*;