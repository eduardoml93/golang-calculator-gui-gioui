package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	for { // Loop infinito
		fmt.Println("Digite o primeiro número:")
		num1Str, _ := reader.ReadString('\n')
		num1Str = strings.TrimSpace(num1Str) // Remove a nova linha do final
		num1, err := strconv.ParseFloat(num1Str, 64)
		if err != nil {
			fmt.Println("Erro ao ler o primeiro número.")
			continue
		}

		fmt.Println("Digite o segundo número:")
		num2Str, _ := reader.ReadString('\n')
		num2Str = strings.TrimSpace(num2Str) // Remove a nova linha do final
		num2, err := strconv.ParseFloat(num2Str, 64)
		if err != nil {
			fmt.Println("Erro ao ler o segundo número.")
			continue
		}

		fmt.Println("Escolha uma operação (+, -, *, /):")
		opStr, _ := reader.ReadString('\n')
		opStr = strings.TrimSpace(opStr) // Remove a nova linha do final

		switch opStr {
		case "+":
			fmt.Printf("%.2f + %.2f = %.2f\n", num1, num2, num1+num2)
		case "-":
			fmt.Printf("%.2f - %.2f = %.2f\n", num1, num2, num1-num2)
		case "*":
			fmt.Printf("%.2f * %.2f = %.2f\n", num1, num2, num1*num2)
		case "/":
			if num2 == 0 {
				fmt.Println("Divisão por zero não é permitida.")
				continue
			}
			fmt.Printf("%.2f / %.2f = %.2f\n", num1, num2, num1/num2)
		default:
			fmt.Println("Operação inválida.")
			continue
		}
	}
}
