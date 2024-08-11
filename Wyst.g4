grammar Wyst;

WS : [ \t\r\n]+ -> skip;

expr: expr ('*'|'/') expr
    | expr ('+'|'-') expr
    | INT
    | '(' expr ')'
    ;

INT: [0-9]+;