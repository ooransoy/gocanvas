package main

import (
	"math/rand"
)

type Orientation uint8

const (
	North Orientation = iota
	East
	South
	West
)

type Rotation int8

const (
	CW  Rotation = 1
	CCW Rotation = -1
)

func RandomOrientation() Orientation {
	return Orientation(rand.Intn(4))
}

func (o Orientation) Rotate(r Rotation) Orientation {
	raw := int8(o) + int8(r)
	if raw >= 0 {
		return Orientation(raw % 4)
	}
start:
	if raw < 0 {
		raw += 4
		goto start
	}
	return Orientation(raw)
}

func (o Orientation) String() string {
	switch o {
	case 0:
		return "North"
	case 1:
		return "East"
	case 2:
		return "South"
	case 3:
		return "West"
	default:
		return "Unknown orientation"
	}
}
