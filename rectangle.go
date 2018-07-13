package main

import (
	"image/color"
	"math"
)

func drawRect(x1, y1, x2 ,y2 int,c color.Color){
	dx := int(math.Abs(float64(x2 - x1)))
	dy := int(math.Abs(float64(y2 - y1)))
	for y := 0;y < dy; y++ {
	  	for x := 0; x < dx; x++ {
	  		canvas.Set(x1 + x, y1 + y, c);
	  	}
	}
}

