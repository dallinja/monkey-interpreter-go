package repl

import (
	"bufio"
	"fmt"
	"io"

	"github.com/dallinja/monkey-interpreter-go/evaluator"
	"github.com/dallinja/monkey-interpreter-go/lexer"
	"github.com/dallinja/monkey-interpreter-go/object"
	"github.com/dallinja/monkey-interpreter-go/parser"
	"github.com/dallinja/monkey-interpreter-go/token"
)

const PROMPT = ">> "

func Start(in io.Reader, out io.Writer) {
	scanner := bufio.NewScanner(in)
	env := object.NewEnvironment()

	for {
		fmt.Fprintf(out, PROMPT)
		scanned := scanner.Scan()
		if !scanned {
			return
		}

		line := scanner.Text()
		l := lexer.New(line)

		// print out lexer results
		for tok := l.NextToken(); tok.Type != token.EOF; tok = l.NextToken() {
			// fmt.Fprintf(out, "%+v\n", tok)
		}

		l2 := lexer.New(line)
		p := parser.New(l2)

		program := p.ParseProgram()
		if len(p.Errors()) != 0 {
			printParserErrors(out, p.Errors())
			continue
		}

		// Print out parser results
		// io.WriteString(out, program.String())
		// io.WriteString(out, "\n")

		evaluated := evaluator.Eval(program, env)
		// print out evaluated results
		if evaluated != nil {
			io.WriteString(out, evaluated.Inspect())
			io.WriteString(out, "\n")
		}
	}
}

const MONKEY_FACE = `            __,__
   .--.  .-"     "-.  .--.
  / .. \/  .-. .-.  \/ .. \
 | |  '|  /   Y   \  |'  | |
 | \   \  \ 0 | 0 /  /   / |
  \ '- ,\.-"""""""-./, -' /
   ''-' /_   ^ ^   _\ '-''
       |  \._   _./  |
       \   \ '~' /   /
        '._ '-=-' _.'
           '-----'
`

func printParserErrors(out io.Writer, errors []string) {
	io.WriteString(out, MONKEY_FACE)
	io.WriteString(out, "Woops! We ran into some monkey business here!\n")
	io.WriteString(out, " parser errors:\n")
	for _, msg := range errors {
		io.WriteString(out, "\t"+msg+"\n")
	}
}
