lexer grammar CommonLexerRules; // lexer grammer 表示这里指存放词法符号规则

MUL: '*';
DIV: '/';
ADD: '+';
SUB: '-';
NUMBER: [0-9]+;
WHITESPACE: [ \r\n\t]+ -> skip;
