package main

import (
	"github.com/faiface/pixel/pixelgl"
)

func main() {
	d := NewDraw()

	pixelgl.Run(d.run)
}