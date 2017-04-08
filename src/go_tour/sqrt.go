package main

import (
	"fmt"
)

func IsClose(x float64, y float64, tolerance float64) bool {
	if (x - y < tolerance) && (y - x < tolerance) {
		return true
	}
	return false
}

func Sqrt(x float64) float64 {
	var z0 float64 = 1.0
	var zn float64
	counter := 1
	for {
		zn = z0 - (z0*z0-x)/2*z0
		if IsClose(z0, zn, 0.001) {
			break
		}
		z0 = zn
		counter += 1
	}

	fmt.Println("Counter: ", counter)
	return zn
}

func main() {
	fmt.Println(Sqrt(2))
}
