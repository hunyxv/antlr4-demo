lexer grammar CommonLexerRules; // lexer grammer 表示这里指存放词法符号规则

MUL: '*';
DIV: '/';
ADD: '+';
SUB: '-';
NUMBER: [0-9]+;
ID : [a-zA-Z]+;
NEWLINE: '\r'? '\n'; // \r 或 \n
WHITESPACE: [ \r\n\t]+ -> skip;
