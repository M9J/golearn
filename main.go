package main

import (
	"fmt"
	"golearn/calculator"
	"golearn/hello"

	"rsc.io/quote"
)

func main() {
	fmt.Println(hello.Greet())
	fmt.Println(quote.Go())
	var a int = calculator.Add(1, 2)
	var b int = calculator.Sub(4, 2)
	var c int = calculator.Mul(5, 6)
	var d int = calculator.Div(9, 3)
	fmt.Println("a", a)
	fmt.Println("b", b)
	fmt.Println("c", c)
	fmt.Println("d", d)
}
