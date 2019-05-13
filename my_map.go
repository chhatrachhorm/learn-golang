package main

import "fmt"

func main() {
	// map[key data type] value data type = shit
	grades := make(map[string]float64)
	grades["john"] = 65.25
	grades["jack"] = 66.32
	grades["todelete"] = 66.32
	delete(grades, "todelete")
	fmt.Println(grades)

	for k, v := range grades {
		fmt.Printf("Key %s = %f\n", k, v)
	}
}
