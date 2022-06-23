package main

import "fmt"

func main() {
	fmt.Println(Sum(4, 4))
}

func Sum(a, b int) int {
	return a + b
}

func Sub(a, b int) int {
	return a - b
}

func Div(a, b int) int {
	return a / b
}

func Mult(a, b int) int {
	return a * b
}