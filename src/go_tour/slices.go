package main

import "golang.org/x/tour/pic"

func Pic(dx, dy int) [][]uint8 {
	picData := make([][]uint8, dy, dy)
	for y := range picData {
		picData[y] = make([]uint8, dx, dx)
		for x:= range picData[y] {
			if isSquare(x, y) {
				picData[y][x] = uint8(255)
			}
		}
	}
	return picData
}

func isLinear(x, y int) bool {
	return y == x
}

func isSquare(x, y int) bool {
	return y == x * x / 250
}

func main() {
	pic.Show(Pic)
}