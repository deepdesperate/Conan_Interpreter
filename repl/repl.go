package repl

import(
	"bufio"
	"fmt"
	"io"
	"github.com/deepdesperate/Conan_Interpreter/lexer"
	"github.com/deepdesperate/Conan_Interpreter/token"
)

const PROMPT = ">>"

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

		for tok:=l.NextToken(); tok.Type != token.EOF; tok = l.NextToken(){
			fmt.Fprintf(out,"%+v\n", tok)
		}
	}
}