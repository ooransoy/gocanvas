package main

import (
	"math"
	"image/color"
)

func drawCircle(originX, originY int, radius float64, color color.Color, fill bool) {
	for dx := -radius; dx <= radius; dx++ {
		for dy := -radius; dy <= radius; dy++ {
			hyp := math.Sqrt(dx*dx+dy*dy)
			if math.Round(hyp) == radius || (fill && hyp < radius) {
				canvas.Set(originX + int(dx), originY + int(dy), color)
			}
		}
	}
}
