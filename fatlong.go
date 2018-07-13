package main

import (
	"image/color"
	"math/rand"
)

type FatLongWalker struct {
	x    int
	y    int
	c    color.Color
	o    Orientation
	wait int
}

func NewFatLongWalker(x, y int) *FatLongWalker {
	w := FatLongWalker{
		x,
		y,
		color.RGBA{rUint8(), rUint8(), rUint8(), 255},
		North,
		0,
	}
	return &w
}

func (w *FatLongWalker) Tick() {
	w.wait++
	if w.wait != 5 {
		return
	}
	w.wait = 0
	canvas.Set(w.x, w.y, w.c)
	canvas.Set(w.x+1, w.y, w.c)
	canvas.Set(w.x-1, w.y, w.c)
	canvas.Set(w.x, w.y+1, w.c)
	canvas.Set(w.x+1, w.y+1, w.c)
	canvas.Set(w.x-1, w.y+1, w.c)
	canvas.Set(w.x, w.y-1, w.c)
	canvas.Set(w.x+1, w.y-1, w.c)
	canvas.Set(w.x-1, w.y-1, w.c)

	switch rand.Intn(10) {
	case 0:
		w.o = w.o.Rotate(CW)
	case 1:
		w.o = w.o.Rotate(CCW)
	}

	switch w.o {
	case North:
		w.y += 3
	case East:
		w.x += 3
	case West:
		w.y -= 3
	case South:
		w.x -= 3
	}
	w.x = ((w.x % width) + width) % width
	w.y = ((w.y % height) + height) % height
}
