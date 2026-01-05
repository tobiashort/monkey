package repl

import (
	"bufio"
	"fmt"
	"io"

	"github.com/tobiashort/monkey/lexer"
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
		l := lexer.New("stdin", line)
		tokens, err := l.Analyze()
		if err != nil {
			fmt.Fprintf(w, "%v\n", err)
		} else {
			for _, t := range tokens {
				fmt.Fprintf(w, "%+v\n", t)
			}
		}
	}
}
