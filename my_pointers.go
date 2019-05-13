package main

import "fmt"

func main() {
	x := 15 //it will try to assign the type (int64)
	a := &x // address of a
	fmt.Println("Address", a)
	fmt.Println("Value", *a)
	*a = 5 + 5
	fmt.Println("updated", *a, x)
}