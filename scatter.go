package gocanvas

import (
	"image/color"
	"math"
)

type ConnectMode uint8

const (
	NoConnect ConnectMode = iota
	OpenConnect
	ClosedConnect
)

func ScatterPlot(d [][2]float64, c color.Color, scale [2]float64, connect ConnectMode) {
	m := func(p [2]float64, axis int) int { // Manipulate
		return []int{Width, Height}[axis]/2 + ftoi(p[axis]*scale[axis])
	}

	var prev = d[len(d)-1]
	if connect == OpenConnect {
		prev = d[0]
	}

	for _, p := range d {
		if connect != NoConnect {
			DrawLine(
				m(p, 0),
				m(p, 1),
				m(prev, 0),
				m(prev, 1),
				c,
			)
			prev = p
		}
		DrawCircle(m(p, 0), m(p, 1), 3, c, true)
	}
}

func ftoi(f float64) int {
	return int(math.Trunc(f))
}
