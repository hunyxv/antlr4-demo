grammar Expr;

import CommonLexerRules;

// Rules
start: expr NEWLINE | ID '=' expr NEWLINE | NEWLINE;

// 优先级和下面定义的顺序有关
expr:
	expr op = (ADD | SUB) expr		# AddSub
	| expr op = (MUL | DIV) expr	# MulDiv
	| NUMBER						# Number
	| ID							# Identity
	| '(' expr ')'					# Parenthesis; // 括号规则
