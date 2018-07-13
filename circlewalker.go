package main

import (
	"image/color"
	"math/rand"
)

type CircleWalker struct {
	x int
	y int
	c color.Color
}

func NewCircleWalker(x, y int) *CircleWalker {
	w := CircleWalker{
		x,
		y,
		color.RGBA{rUint8(), rUint8(), rUint8(), 255},
	}
	return &w
}

func (w *CircleWalker) Tick() {
	drawCircle(w.x, w.y, 10, w.c, false)
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
