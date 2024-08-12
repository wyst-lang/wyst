grammar Wyst;

WS : [ \t\r\n]+ -> skip;

expr: IDENTIFIER
    ;

INT: [0-9]+;
IDENTIFIER: [a-zA-Z_] ([a-zA-Z0-9] | '::')*;