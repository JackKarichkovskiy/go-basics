package main

import (
	"fmt"
)

func Pic(dx, dy int) [][]uint8 {
	pic := make([][]uint8, dy)
	for i := range pic {
		pic[i] = make([]uint8, dx)
	}
	return pic
}

func PicMain() {
	pic := Pic(8, 6)
	fmt.Println(pic)
}
