package main

import (
	"image/color"
	"math"
)

func drawLine(x1, y1, x2 ,y2 int,c color.RGBA){
	dx := int(math.Abs(float64(x2 - x1)))
	dy := int(math.Abs(float64(y2 - y1)))
	for i := 0;i < dx; i++ {
	  	canvas.Set(x1+i,y1+(dy*(x1+i - x1))/dx,c)
	}
}

