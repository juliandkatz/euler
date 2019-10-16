package main

import "fmt"

func simple() int {
	sum := 0

	for x := 3; x < 1000; x++ {
		if x%3 == 0 || x%5 == 0 {
			sum += x
		}
	}

	return sum
}

func main() {
	fmt.Println(simple())
}
