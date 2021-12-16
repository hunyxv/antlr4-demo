grammar Rows;

@parser::members {    // add members to generated RowsParser

var col int
func NewMyRowsParser(input antlr.TokenStream, c int) *RowsParser {
	col = c
	return NewRowsParser(input)
}
}

file: (row NL)+;

row
	locals[int i=0]: (
		STUFF {
            $i++;
            if ( $i == col ) {
                fmt.Println($STUFF.text)
            }
        }
	)+;

TAB: '\t' -> skip; // match but don't pass to the parser
NL: '\r'? '\n'; // match and pass to the parser
STUFF: ~[\t\r\n]+; // match any chars except tab, newline