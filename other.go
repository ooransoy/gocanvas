package main

import (
	"image"
)

func resetCanvas() {
	canvas = image.NewRGBA(image.Rect(0, 0, width, height))
}
