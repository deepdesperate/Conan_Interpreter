package repl

import(
	"bufio"
	"fmt"
	"io"
	"github.com/deepdesperate/Conan_Interpreter/lexer"
	"github.com/deepdesperate/Conan_Interpreter/parser"
	"github.com/deepdesperate/Conan_Interpreter/evaluator"

)

const PROMPT = ">>"
const CONAN_FACE = `
"--\__/-\__/--"
`


func Start(in io.Reader, out io.Writer){
	scanner := bufio.NewScanner(in)
	for {
		fmt.Fprintf(out, PROMPT)
		scanned := scanner.Scan()

		if !scanned{
			return
		}

		line := scanner.Text()
		l:=lexer.New(line)
		p:=parser.New(l)

		program := p.ParseProgram()
		if len(p.Errors()) != 0{
			printParseErrors(out, p.Errors() )
			continue
		}

		evaulated := evaluator.Eval(program)
		if evaulated != nil {
			io.WriteString(out, evaulated.Inspect())
			io.WriteString(out,"\n")
		}	
		
	}
}

func printParseErrors(out io.Writer, errors []string){
	io.WriteString(out, CONAN_FACE)
	io.WriteString(out, "Whoops! We ran into some mystery here! \n")
	io.WriteString(out,"parse erros: \n")
	for _,msg := range errors{
		io.WriteString(out, "\t"+msg+"\n")
	}
}