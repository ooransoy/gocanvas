package gocanvas

import (
	"image"
	"math/rand"
	"time"
)

// RNG Seeder
var _ = func() struct{} { rand.Seed(time.Now().UnixNano()); return struct{}{} }()

var Width = 800
var Height = 600
var Tick = time.Millisecond

var Canvas *image.RGBA = image.NewRGBA(
	image.Rect(0, 0, Width, Height),
)

var Loop = func() {
	panic("No loop function set")
}

func Start() {
	go serve()
	for {
		Loop()
		time.Sleep(Tick)
	}
}
