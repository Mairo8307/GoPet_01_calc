package main

import (
	"fmt"
)

func main() {
	var n1, n2 float32
	var action string
	fmt.Scan(&n1, &n2, &action)
	switch action {
	case "+":
		fmt.Println(n1 + n2)
	case "-":
		fmt.Println(n1 - n2)
	case "*":
		fmt.Println(n1 * n2)
	case "%":
		fmt.Println(int(n1) % int(n2))
	case "/":
		if n2 != 0 {
			fmt.Println(n1 / n2)
		} else {
			fmt.Println("Делить на ноль нельзя!")
		}
	}
}
