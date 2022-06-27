package calculator

import "golearn/calculator/arithmetic"

func Add(x int, y int) int {
	return arithmetic.Addition(x, y)
}

func Sub(x int, y int) int {
	return arithmetic.Subtraction(x, y)
}

func Mul(x int, y int) int {
	return arithmetic.Multiplication(x, y)
}

func Div(x int, y int) int {
	return arithmetic.Division(x, y)
}
