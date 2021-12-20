grammar Data;

file: group+;

group: INT sequence[$INT.int];

sequence[int n]
	locals[int i=0]: ({$i<$n}? INT {$i++;})*;

INT: [0-9]+;
WS: [ \t\n\r]+ -> skip;