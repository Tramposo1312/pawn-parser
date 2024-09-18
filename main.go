package main

import (
	"fmt"
	"io/ioutil"
	"os"

	"github.com/Tramposo1312/pawn-parser/ast"
	"github.com/Tramposo1312/pawn-parser/lexer"
	"github.com/Tramposo1312/pawn-parser/parser"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: go run main.go <filename.pwn>")
		os.Exit(1)
	}

	filename := os.Args[1]
	content, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Printf("Error reading file: %v\n", err)
		os.Exit(1)
	}

	l := lexer.New(string(content))
	p := parser.New(l)

	program, err := p.ParseProgram()
	if err != nil {
		fmt.Printf("Parser errors:\n%v\n", err)
		os.Exit(1)
	}

	fmt.Println("Parsing completed successfully.")
	printer := ast.NewAstPrinter()
	fmt.Println(printer.Print(program))
}
