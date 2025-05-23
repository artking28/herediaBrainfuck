package root

import (
	"errors"
	"fmt"
	"math"
)

const DEFAULT_SIZE = 30_000

type BfTuring struct {
	tape   []byte
	cursor uint
}

func NewBfTuring() BfTuring {
	return BfTuring{make([]byte, DEFAULT_SIZE), 0}
}

func NewBfTuringAll(tape []byte, cursor uint) BfTuring {
	size := DEFAULT_SIZE - len(tape)
	if len(tape) > DEFAULT_SIZE {
		i := int(math.Ceil(float64(len(tape)) / DEFAULT_SIZE))
		size = i*DEFAULT_SIZE - len(tape)
	}
	return BfTuring{
		append(tape, make([]byte, size)...), 
		cursor,
	}
}

func (this *BfTuring) Execute(input string) error {
	for i := 0; i < len(input); i++ {
		switch input[i] {
		case '>':
			this.cursor++
			break
		case '<':
			this.cursor--
			break
		case '+':
			if this.cursor >= uint(len(this.tape)) {
				this.tape = append(this.tape, make([]byte, 1024)...)
			}
			this.tape[this.cursor]++
			break
		case '-':
			if this.cursor < 0 {
				return errors.New("index out of bounds, substracting cell at position -1.")
			}
			this.tape[this.cursor]--
			break
		case '.':
			print(string(this.tape[this.cursor]))
		case ',':
			var ch rune
			println("Please, type a ASCII character:")
			if _, err := fmt.Scanf("%c", &ch); err != nil {
				return err
			}
			if ch > 127 {
				return errors.New("It ain't a ASCII value.")
			}
			this.tape[this.cursor] = byte(ch)
			break
		case '[':
			if this.tape[this.cursor] != 0 {
				break
			}
			i += 1
			for mem := 0; true; i++ {
				if i >= len(input) {
					return errors.New("index out of bounds, reading program beyond its limits.")
				}
				if input[i] == '[' {
					mem++
				}
				if input[i] == ']' {
					if mem == 0 {
						break
					}
					mem--
				}
			}
			break
		case ']':
			if this.tape[this.cursor] == 0 {
				break
			}
			i -= 1
			for mem := 0; true; i-- {
				if i < 0 {
					return errors.New("index out of bounds, reading program at position -1.")
				}
				if input[i] == ']' {
					mem++
				}
				if input[i] == '[' {
					if mem == 0 {
						break
					}
					mem--
				}
			}
			break
		default:
			break
			// return fmt.Errorf("unkown character '%c' at position %d.", input[i], i)
		}
	}
	return nil
}
