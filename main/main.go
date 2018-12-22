package main

import (
	"fmt"
	"image/color"
	"math/rand"
	"time"

	gc "github.com/ooransoy/gocanvas"
)

func init() {
	fmt.Println("Starting GoCanvas")
	var seq = [][2]float64{
		{0.1, -2.6},
		{-4, -3},
		{-1.5, -0.75},
		{0, 0},
		{-2.25, 2},
		{1.1, 2.5},
		{3.4, 1.5},
		{2.65, -1.75},
		{1, -1.2},
	}
	c := color.RGBA{
		uint8(rand.Intn(256)),
		uint8(rand.Intn(256)),
		uint8(rand.Intn(256)),
		255,
	}

	gc.Width = 800
	gc.Height = 600
	gc.Tick = time.Millisecond
	gc.Loop = func() {
		gc.ScatterPlot(seq, c, [2]float64{50, -50}, gc.ClosedConnect)
	}
}

func main() {
	fmt.Println("CTRL-C to exit.")
	gc.Start()
}
