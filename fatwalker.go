package main

import (
	"image/color"
)

type FatWalker struct {
	x int
	y int
	c color.Color
	wait int
}

func NewFatWalker(x, y int) *FatWalker {
	w := FatWalker{
		x,
		y,
		color.RGBA{rUint8(), rUint8(), rUint8(), 255},
		0,
	}
	return &w
}

func (w *FatWalker) Tick() {
	w.wait++
	if w.wait != 5 {
		return
	}
	w.wait = 0
	canvas.Set(w.x  , w.y  , w.c)
	canvas.Set(w.x+1, w.y  , w.c)
	canvas.Set(w.x-1, w.y  , w.c)
	canvas.Set(w.x  , w.y+1, w.c)
	canvas.Set(w.x+1, w.y+1, w.c)
	canvas.Set(w.x-1, w.y+1, w.c)
	canvas.Set(w.x  , w.y-1, w.c)
	canvas.Set(w.x+1, w.y-1, w.c)
	canvas.Set(w.x-1, w.y-1, w.c)

	switch RandomOrientation() {
	case North:
		w.y+=3
	case East:
		w.x+=3
	case West:
		w.y-=3
	case South:
		w.x-=3
	}

	w.x = ((w.x % width) + width) % width
	w.y = ((w.y % height) + height) % height
}
