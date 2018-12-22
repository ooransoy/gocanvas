package gocanvas

import (
	"image"
	"image/color"
	"math"
)

func DrawLine(x1, y1, x2, y2 int, c color.Color) {
	dx := x2 - x1
	dy := y2 - y1

	if int(math.Abs(float64(dx))) >= int(math.Abs(float64(dy))) {
		for i := float64(0); ftoi(i) != dx; {
			Canvas.Set(
				x1+ftoi(i),
				y1+ftoi(float64(dy)/float64(dx)*i),
				c,
			)

			if dx < 0 {
				i -= 0.01
				continue
			}
			i += 0.01
		}
	} else {
		for i := float64(0); ftoi(i) != dy; {
			Canvas.Set(
				x1+ftoi(float64(dx)/float64(dy)*i),
				y1+ftoi(i),
				c,
			)

			if dy < 0 {
				i -= 0.01
				continue
			}
			i += 0.01
		}
	}
}

func DrawCircle(originX, originY int, radius float64, c color.Color, fill bool) {
	for dx := -radius; dx <= radius; dx++ {
		for dy := -radius; dy <= radius; dy++ {
			hyp := math.Sqrt(dx*dx + dy*dy)
			if math.Round(hyp) == radius || (fill && hyp < radius) {
				Canvas.Set(originX+int(dx), originY+int(dy), c)
			}
		}
	}
}

func DrawRect(x1, y1, x2, y2 int, c color.Color) {
	dx := int(math.Abs(float64(x2 - x1)))
	dy := int(math.Abs(float64(y2 - y1)))
	for y := 0; y < dy; y++ {
		for x := 0; x < dx; x++ {
			Canvas.Set(x1+x, y1+y, c)
		}
	}
}

func ResetCanvas() {
	Canvas = image.NewRGBA(image.Rect(0, 0, Width, Height))
}

/*
func intabs(n int) int {
	return int(math.Abs(float64(n)))
}*/
