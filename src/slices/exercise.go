package main

import "golang.org/x/tour/pic"

func Pic(dx, dy int) [][]uint8 {
	image := make([][]uint8, dy)
	for i := 0; i < dy; i++ {
		image[i] = make([]uint8, dx)
		for k := 0; k < dx; k++ {
			image[i][k] = uint8((k+1)%(i+1))
		}
	}
	return image
}

func main() {
	pic.Show(Pic)
}
