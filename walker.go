package main

import (
	"image/color"
	"math/rand"
)

type Walker struct {
	x int
	y int
	c color.Color
}

func NewWalker(x, y int) *Walker {
	w := Walker{
		x,
		y,
		color.RGBA{rUint8(), rUint8(), rUint8(), 255},
	}
	return &w
}

func (w *Walker) Tick() {
	canvas.Set(w.x, w.y, w.c)
	if rand.Intn(10) == 1 {
		canvas.Set(w.x, w.y, color.RGBA{0, 0, 0, 255})
	}
	switch RandomOrientation() {
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

func rUint8() uint8 {
	return uint8(rand.Intn(256))
}
