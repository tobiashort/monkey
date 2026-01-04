package repl

import (
	"bufio"
	"fmt"
	"io"

	"github.com/tobiashort/monkey/lexer"
	"github.com/tobiashort/monkey/token"
)

const PROMPT = ">> "

func Start(w io.Writer, r io.Reader) {
	scanner := bufio.NewScanner(r)

	for {
		fmt.Fprintf(w, PROMPT)

		scanned := scanner.Scan()
		if !scanned {
			return
		}

		line := scanner.Text()
		l := lexer.New(line)

		for tok := l.NextToken(); tok.Type != token.EOF; tok = l.NextToken() {
			fmt.Fprintf(w, "%+v\n", tok)
		}
	}
}
