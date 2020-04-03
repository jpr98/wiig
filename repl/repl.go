package repl

import (
	"bufio"
	"fmt"
	"io"

	"github.com/jpr98/wiig/lexer"
	"github.com/jpr98/wiig/token"
)

const PROMPT = ">>"

// Start initiates the code execution loop
func Start(in io.Reader, out io.Writer) {
	scanner := bufio.NewScanner(in)

	for {
		fmt.Print(PROMPT)
		scanned := scanner.Scan()
		if !scanned {
			return
		}

		line := scanner.Text()
		if line == "!!exit" {
			break
		}
		l := lexer.New(line)

		for tok := l.NextToken(); tok.Type != token.EOF; tok = l.NextToken() {
			fmt.Printf("%+v\n", tok)
		}
	}
}
