package main

import "fmt"

func add(num1 float64, num2 float64) float64 {
	return num1 + num2
}

func echo(s1 string, s2 string) (string, string)  {
	return s1, s2
}

func main() {
	var x float64 = 5
	var y float64 = 6.3
	fmt.Println("Sum of x and y =", add(x, y))

	fmt.Println(echo("Test", "ECHO"))
}
