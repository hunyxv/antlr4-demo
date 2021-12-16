grammar Hello;

import CommonLexerRules;
// // Tokens
// MUL: '*';
// DIV: '/';
// ADD: '+';
// SUB: '-';
// NUMBER: [0-9]+;
// WHITESPACE: [ \r\n\t]+ -> skip;

// Rules
start: expr EOF;

// 优先级和下面定义的顺序有关
expr:
	expr op = ('*' | '/') expr		# MulDiv
	| expr op = ('+' | '-') expr	# AddSub
	| NUMBER						# Number
	| '(' expr ')'					# Parenthesis; // 括号规则
