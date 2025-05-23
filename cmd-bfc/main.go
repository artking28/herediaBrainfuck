package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

// Token representa um número ou operador
type Token struct {
	isOperator bool
	value      string
}

func printfuncString() string {
	return ">++++++++++<<[->+>-[>+>>]>[+[-<+>]>+>>]<<<<<<]>>[-]>>>++++++++++<[->-[>+>>]>[+[-<+>]>+>>]<<<<<]>[-]>>[>++++++[-<++++++++>]<.<<+>+>[-]]<[<[->-<]++++++[->++++++++<]>.[-]]<<++++++[-<++++++++>]<.[-]<<[-<+>]"
}

// Função para tokenizar a expressão
func tokenize(expr string) ([]Token, error) {
	expr = strings.ReplaceAll(expr, " ", "")
	var tokens []Token
	numBuf := strings.Builder{}
	for _, c := range expr {
		if c >= '0' && c <= '9' {
			numBuf.WriteRune(c)
		} else if c == '+' || c == '-' || c == '*' {
			if numBuf.Len() > 0 {
				tokens = append(tokens, Token{false, numBuf.String()})
				numBuf.Reset()
			}
			tokens = append(tokens, Token{true, string(c)})
		} else {
			return nil, fmt.Errorf("caractere inválido: %c", c)
		}
	}
	if numBuf.Len() > 0 {
		tokens = append(tokens, Token{false, numBuf.String()})
	}
	return tokens, nil
}

// Função para converter tokens em notação pós-fixa (usando Shunting Yard)
func toPostfix(tokens []Token) ([]Token, error) {
	var output []Token
	var operators []Token
	for _, token := range tokens {
		if !token.isOperator {
			output = append(output, token)
		} else {
			for len(operators) > 0 {
				top := operators[len(operators)-1]
				if (token.value == "+" || token.value == "-") && (top.value == "*" || top.value == "+") {
					output = append(output, top)
					operators = operators[:len(operators)-1]
				} else {
					break
				}
			}
			operators = append(operators, token)
		}
	}
	for len(operators) > 0 {
		output = append(output, operators[len(operators)-1])
		operators = operators[:len(operators)-1]
	}
	return output, nil
}

// Função para gerar Brainfuck para imprimir uma string (nome + "=")
func generateBrainfuckPrintString(s string) string {
	var bf strings.Builder
	for _, c := range s {
		bf.WriteString("[-]")
		for i := 0; i < int(c); i++ {
			bf.WriteString("+")
		}
		bf.WriteString(".>")
	}
	return bf.String()
}

// Função para gerar código Brainfuck a partir de tokens em pós-fixa
func generateBrainfuck(postfix []Token) (string, error) {
	var bfCode strings.Builder
	cellIndex := 0

	for _, token := range postfix {
		if !token.isOperator {
			num, err := strconv.Atoi(token.value)
			if err != nil {
				return "", fmt.Errorf("número inválido: %s", token.value)
			}

			bfCode.WriteString("[-]")
			// Incrementa valor
			for i := 0; i < num; i++ {
				bfCode.WriteString("+")
			}
			bfCode.WriteString(">")
			cellIndex++
		} else {
			if cellIndex < 2 {
				return "", fmt.Errorf("expressão inválida: operandos insuficientes")
			}

			bfCode.WriteString(strings.Repeat("<", 2))
			switch token.value {
			case "+":
				// [->+<] move valor da célula da direita pra esquerda (soma)
				bfCode.WriteString("[->+<]>[-<+>]<")
			case "-":
				// Subtrai célula da direita da da esquerda: [->-<]
				bfCode.WriteString("[->-<]>[-<+>]<")
			case "*":
				// Multiplica célula da esquerda pela da direita
				// usa célula extra temporária
				bfCode.WriteString(">>[-]<<[>[->+>+<<]>>[-<<+>>]<<<-]>>[-<<+>>]<<")
			default:
				return "", fmt.Errorf("operador inválido: %s", token.value)
			}

			bfCode.WriteString(">")
			cellIndex--
		}
	}

	bfCode.WriteString(printfuncString())
	return bfCode.String(), nil
}

func main() {
	expr := strings.Join(os.Args[1:], "")
	expr = strings.ReplaceAll(expr, "\"", "")

	split := strings.Split(expr, "=")
	if len(split) != 2 {
		println("error, siga o padrao '<nome> = <expressao>'")
		os.Exit(1)
	}

	name, value := strings.TrimSpace(split[0]), split[1]

	// Tokeniza
	tokens, err := tokenize(value)
	if err != nil {
		fmt.Printf("Erro ao tokenizar %s: %v\n", value, err)
		os.Exit(1)
	}

	// Converte para pós-fixa
	postfix, err := toPostfix(tokens)
	if err != nil {
		fmt.Printf("Erro ao converter %s: %v\n", value, err)
		os.Exit(1)
	}

	// Gera Brainfuck para imprimir nome + "="
	namePrint := generateBrainfuckPrintString(name + "=")

	// Gera Brainfuck para calcular e imprimir resultado
	bfCode, err := generateBrainfuck(postfix)
	if err != nil {
		fmt.Printf("Erro ao gerar Brainfuck para %s: %v\n", value, err)
		os.Exit(1)
	}

	// Junta os dois códigos
	fmt.Printf(namePrint + bfCode)
}
