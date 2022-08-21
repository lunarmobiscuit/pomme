package pomme

import (
	"fmt"
)

// Linked list of symbols
type keyword struct {
	name	string		// the "keyword"
	f		kFund		// the func that parses the keyword
}
type kFund func (*parser) error

var keywords = []keyword {
	{"print", parsePrint},
	{"os", parseOs},
	{"if", parseIf},
	{"for", parseFor},
	{"while", parseWhile},
	{"var", parseVar},
}


/*
 *  Lookup the keyword
 */
func (p *parser) isKeyword(token string) bool {
	// Find the table entry for the keyword
	for k := range keywords {
		if (token == keywords[k].name) {
			return true
		}
	}

	return false
}

/*
 *  Lookup and parse the keyword
 */
func (p *parser) parseKeyword(token string) error {
	// Find the parser for the keyword
	for k := range keywords {
		if (token == keywords[k].name) {
			return keywords[k].f(p)
		}
	}

	return fmt.Errorf("keyword '%s' is invalid", token)
}

/*
 *  Parse the 'print' keyword
 */
func parsePrint(p *parser) error {
	fmt.Printf("PRINT is not yet supported [%d-%d]\n", p.i, p.n)
	return nil
}

/*
 *  Parse the 'os' keyword
 */
func parseOs(p *parser) error {
	fmt.Printf("IF is not yet supported [%d-%d]\n", p.i, p.n)
	return nil
}

/*
 *  Parse the 'if' keyword
 */
func parseIf(p *parser) error {
	fmt.Printf("IF is not yet supported [%d-%d]\n", p.i, p.n)
	return nil
}

/*
 *  Parse the 'for' keyword
 */
func parseFor(p *parser) error {
	fmt.Printf("FOR is not yet supported [%d-%d]\n", p.i, p.n)
	return nil
}

/*
 *  Parse the 'while' keyword
 */
func parseWhile(p *parser) error {
	fmt.Printf("WHILE is not yet supported [%d-%d]\n", p.i, p.n)
	return nil
}

/*
 *  Parse the 'var' keyword
 */
func parseVar(p *parser) error {
	fmt.Printf("VAR is not yet supported [%d-%d]\n", p.i, p.n)
	return nil
}
