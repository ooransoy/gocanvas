package main

import (
	"fmt"
	"image"
)

func getImgUrl(n int) string {
	return fmt.Sprintf("./images/img%06d.png", n)
}

func resetCanvas() {
	canvas = image.NewRGBA(image.Rect(0, 0, width, height))
}
