package main

import "fmt"

func main() {
	for i := 0; i<10; i++ {
		fmt.Printf("i value is %d", i)
	}
	i := 0
	for {
		fmt.Println("I am infinite for loop")
		if i > 10{
			break
		}
		i++
	}
}