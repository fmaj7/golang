package main

import "fmt"

// fibonacci is a function that returns
// a function that returns an int.
func fibonacci() func() int {
	n := 0
	nPlus1 := 1
	return func() int {
		output := n
		n = nPlus1
		nPlus1 = n + output
		return output
	}
}

func main() {
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}
