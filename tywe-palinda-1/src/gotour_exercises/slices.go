package main

import "golang.org/x/tour/pic"

func Pic(dx, dy int) [][]uint8 {
	picture := make([][]uint8, dy)
	numbs := make([]uint8, dx)

	for i := 0; i < dy; i++ {
		for j := 0; j < dx; j++ {
			numbs[j] = uint8(i ^ j)
		}
		picture[i] = numbs
	}
	return picture
}

func main() {
	pic.Show(Pic)
}
