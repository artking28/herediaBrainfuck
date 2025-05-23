package main

import (
	"io"
	"os"
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
	tm := NewBfTuring()
	if err := tm.Execute(input); err != nil {
		panic(err.Error())
	}
}
