package main

import (
	"os"
	"io"
	"p3-brainfuck"
)

func main() {
	content, err := io.ReadAll(os.Stdin)
	if err != nil {
		panic(err.Error())
	}
	if len(content) <= 0 {
		println("empty program")
		os.Exit(0)
	}
	
	input := string(content)
	tm := root.NewBfTuring()
	if err := tm.Execute(input); err != nil {
		panic(err.Error())
	}
}
