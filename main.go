package main

import (
	"fmt"
	"image"
	"image/color"
	"math/rand"
	"time"
)

// RNG Seeder
var _ = func() struct{} { rand.Seed(time.Now().UnixNano()); return struct{}{} }()

// Dimensions of screen
const width int = 800
const height int = 600
const tickWait time.Duration = 1 * time.Millisecond

var canvas *image.RGBA = image.NewRGBA(
	image.Rect(0, 0, width, height),
)

func main() {
	fmt.Println("Starting game...")
	// Setup
	cWalkerSlice := make([]*CircleWalker, 10, 10)
	for i, _ := range cWalkerSlice {
		cWalkerSlice[i] = NewCircleWalker(width/2, height/2)
	}
	// /Setup

	go serve()
	fmt.Println("Server is up! CTRL-C to exit")

	// Game Loop
	for {
		if rand.Int()%2000 == 0 {
			col := pickcol(cWalkerSlice)
			for i := 0; i < 800; i++ {
				for j := 0; j < 600; j++ {
					canvas.Set(i, j, col)
				}
			}
		}
		for _, w := range cWalkerSlice {
			w.Tick()
		}
		time.Sleep(tickWait)
	}
	// /Game Loop
}

func pickcol(a []*CircleWalker) color.Color {
	return (*a[rand.Intn(len(a))]).c
}
