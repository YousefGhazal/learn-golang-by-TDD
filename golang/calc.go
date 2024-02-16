package main

import "fmt"

func main() {

	var opretor string
	var num1 int
	var num2 int
	fmt.Scan(&num1)
	fmt.Scan(&opretor)
	fmt.Scan(&num2)

	calc := make(map[string]func(int, int))

	calc["+"] = func(num1 int, num2 int) {
		sum := num1 + num2
		fmt.Println(sum)
	}

	calc["-"] = func(num1 int, num2 int) {
		sub := num1 - num2
		fmt.Println(sub)
	}

	calc["*"] = func(num1 int, num2 int) {
		mul := num1 * num2
		fmt.Println(mul)
	}

	calc["/"] = func(num1 int, num2 int) {
		div := num1 / num2
		fmt.Println(div)
	}

	calc[opretor](num1, num1)

}
