package main

import (
	"image/color"
	"math/rand"
)

type LongWalker struct {
	x int
	y int
	c color.Color
	o Orientation
}

func NewLongWalker(x, y int) *LongWalker {
	w := LongWalker{
		x,
		y,
		color.RGBA{rUint8(), rUint8(), rUint8(), 255},
		North,
	}
	return &w
}

func (w *LongWalker) Tick() {
	canvas.Set(w.x, w.y, w.c)
	switch rand.Intn(10) {
	case 0:
		w.o = w.o.Rotate(CW)
	case 1:
		w.o = w.o.Rotate(CCW)
	}

	switch w.o {
	case North:
		w.y++
	case East:
		w.x++
	case West:
		w.y--
	case South:
		w.x--
	}
	w.x = ((w.x % width) + width) % width
	w.y = ((w.y % height) + height) % height
}
