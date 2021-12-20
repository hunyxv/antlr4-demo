lexer grammar XMLLexer;

// ------ 默认模式
OPEN: '<' -> pushMode(INSIDE); // 模式切换
COMMENT: '<!--' .*? '-->'; // 注释
EntityRef: '&' [a-z]+ ';'; // 实体引用，比如 &lt; &gt; &amp; 
TEXT: ~('<' | '&')+;

// ------ 标签之中的东西
mode INSIDE;
CLOSE: '>' -> popMode; // 返回到默认模式
SLASH_CLOSE: '/>' -> popMode; // 返回默认模式
EQUALS: '=';
STRING: '"' .*? '"';
SlashName: '/' Name;
Name: ALPHA (ALPHA | DIGIT)*;
S: [ \t\n\r] -> skip;

fragment ALPHA: [a-zA-Z];

fragment DIGIT: [0-9];