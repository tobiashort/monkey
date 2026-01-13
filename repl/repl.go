package repl

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"os"

	"github.com/tobiashort/monkey/lexer"
	"github.com/tobiashort/monkey/parser"
	"github.com/tobiashort/utils-go/must"
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
			os.Exit(1)
		}

		p := parser.New(tokens)
		ast, err := p.Parse()
		if err != nil {
			fmt.Fprintf(w, "%v\n", err)
		} else {
			j := string(must.Do2(json.MarshalIndent(ast, "", "  ")))
			fmt.Fprintf(w, "%s\n", j)
		}
	}
}
