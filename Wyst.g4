grammar Wyst;

WS: [ \t\r\n]+ -> skip;
NUMBER: [1-9][0-9]* | '0';
HEX: '0x' [a-fA-F0-9]*;
IDENTIFIER: [a-zA-Z_] [a-zA-Z0-9_]* ('::' [a-zA-Z_] [a-zA-Z0-9_]*)*;
GOIDENTIFIER: '^' [a-zA-Z_] [a-zA-Z0-9_]* ('.' [a-zA-Z_] [a-zA-Z0-9_]*)*;
MATH: '+' | '-' | '*' | '/' | '==' | '!=' | '>=' | '<=';
fragment ESC: '\\' ['"\\] ;
STRING: '"' (ESC | ~["\\])* '"';
MODULE_NAME: [a-zA-Z_] [a-zA-Z0-9_]*;


round_def: '(' (var_def (',' var_def)* ','?)? ')';
enum_curly: '{' (var_def ';')* '}';
round_call: '(' (expr (',' expr)* ','?)? ')';
fn_call: IDENTIFIER round_call;
func_def: IDENTIFIER IDENTIFIER round_def code_block;
var_def: IDENTIFIER IDENTIFIER;
var_def_set: var_def '=' expr;
call_tree: (fn_call|IDENTIFIER) ('.' (fn_call|IDENTIFIER))*;
struct_def: 'struct' IDENTIFIER enum_curly;
namespace: 'namespace' IDENTIFIER '{' (((var_def|var_def_set) ';')|func_def)* '}';
import_statement: 'include' MODULE_NAME ';';
use_statement: 'use' IDENTIFIER ';';
go_import: '%import' STRING ';';
map: '%map' IDENTIFIER ',' IDENTIFIER ';';
go_call: GOIDENTIFIER round_call;

// if else statements
if_expr: '(' expr (('&&'|'||') expr)* ')';
if_statement: 'if' if_expr code_block*;
elseif_statement: 'else' if_statement;
else_statement: 'else' code_block*;
if_tree: if_statement elseif_statement* else_statement?;

expr:
    (
        call_tree |
        NUMBER |
        HEX |
        IDENTIFIER |
        STRING |
        go_call |
        GOIDENTIFIER
    ) (MATH expr)*;

code_block: '{' (((
    expr |
    var_def |
    var_def_set
) ';') | (
    if_tree
))* '}';

top: (
    func_def |
    struct_def |
    namespace |
    import_statement |
    use_statement |
    go_import |
    map
)*;