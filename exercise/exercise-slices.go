package main

import "golang.org/x/tour/pic"

func Pic(dx, dy int) [][]uint8 {
	v := make([][]uint8, dy, dx)
	for i := 0; i < dy; i++ {
		row := make([]uint8, dx)
		for j := 0; j < dx; j++ {
			row[j] = uint8((i + j) / 2)
			//row[j] = uint8(i * j)
			//row[j] = uint8(i ^ j)
		}
		v[i] = row
	}
	return v
}

func main() {
	pic.Show(Pic)
}
