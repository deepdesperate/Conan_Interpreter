package repl

import(
	"bufio"
	"fmt"
	"io"
	"github.com/deepdesperate/Conan_Interpreter/compiler"
	"github.com/deepdesperate/Conan_Interpreter/lexer"
	"github.com/deepdesperate/Conan_Interpreter/parser"
	"github.com/deepdesperate/Conan_Interpreter/vm"

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

		comp := compiler.New()
		err := comp.Compile(program)

		if err != nil {
			fmt.Fprintf(out, "Woops! Compilation failed:\n %s\n", err)
			continue
		}

		machine := vm.New(comp.Bytecode())
		err = machine.Run()
		if err != nil {
			fmt.Fprintf(out, "Woops! Executing bytecode failed:\n %s\n", err)
			continue
		}

		stackTop := machine.StackTop()
		
		io.WriteString(out, stackTop.Inspect())
		io.WriteString(out, "\n")

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