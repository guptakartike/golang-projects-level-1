package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func add(a int, b int) int {
	return a + b
}
func sub(a int, b int) int {
	return a - b
}
func prod(a int, b int) int {
	return a * b
}
func div(a int, b int) int {

	return a / b
}
func mod(a int, b int) int {

	return a % b
}

func main() {

	fmt.Println("!! Welcome to Kartike's Calculator !!")
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Println("Calculator is running...")

		fmt.Print("Enter first number: ")
		scanner.Scan()
		aStr := strings.TrimSpace(scanner.Text())
		a, err := strconv.Atoi(aStr)
		if err != nil {
			fmt.Println("INVALID INPUT")
			continue
		}

		fmt.Print("Enter second number: ")
		scanner.Scan()
		bStr := strings.TrimSpace(scanner.Text())
		b, err := strconv.Atoi(bStr)
		if err != nil {
			fmt.Println("INVALID INPUT")
			continue
		}

		fmt.Print("Enter the operation(+,-,*,/,%): ")
		scanner.Scan()
		op := strings.TrimSpace(scanner.Text())

		switch op {
		case "+":
			fmt.Println(add(a, b))
		case "-":
			fmt.Println(sub(a, b))
		case "*":
			fmt.Println(prod(a, b))
		case "/":
			if b == 0 {
				fmt.Println("Error: Division by zero")
				continue
			}
			fmt.Println(div(a, b))
		case "%":
			if b == 0 {
				fmt.Println("Error: Division by zero")
				continue
			}
			fmt.Println(mod(a, b))
		default:
			fmt.Println("Invalid operation")
			continue
		}

		fmt.Println("Do you want to continue? (Y/N)")
		scanner.Scan()
		choice := strings.TrimSpace(scanner.Text())
		if choice != "Y" && choice != "y" {
			break
		}
	}

}
