package main

import (
	"github.com/faiface/pixel/pixelgl"
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/imdraw"
	"golang.org/x/image/colornames"
)

type D struct {
	Size   int
	Cfg    *pixelgl.WindowConfig
	Imd    *imdraw.IMDraw
	Win    *pixelgl.Window
	Pad    float64
	Offset pixel.Vec
}

func (d *D) run() {
	d.Init()

	imd := d.Imd

	d.DrawBoard()
	for !d.Win.Closed() {
		imd.Draw(d.Win)
		d.Win.Update()
	}
}

func NewDraw() *D {
	d := &D{}
	return d
}

func (d *D) Init() {
	var err error
	d.Size = 7
	d.Pad = 40.0
	d.Cfg = &pixelgl.WindowConfig{
		Title: "Smallest sum of products - Grid",
		Bounds: pixel.R(0, 0, float64(d.Size)*d.Pad+d.Pad*2,
			float64(d.Size)*d.Pad+d.Pad*5),
		VSync: true,
	}

	w := d.Pad * float64(d.Size)
	offset := pixel.V(d.Cfg.Bounds.Center().X-w/2,
		d.Cfg.Bounds.Size().Y-float64(d.Size+1)*d.Pad)
	d.Offset = offset

	d.Win, err = pixelgl.NewWindow(*d.Cfg)
	d.Win.Clear(colornames.Aliceblue)


	if err != nil {
		panic(err)
	}
	d.Imd = imdraw.New(nil)
}

func (d *D) DrawBoard(reset ...bool) {
	d.Imd.Color = colornames.Black
	d.Imd.EndShape = imdraw.RoundEndShape
	d.Imd.Clear()
	xypad := pixel.V(d.Pad, d.Pad)
	for i := 0; i < d.Size; i++ {
		for j := 0; j < d.Size; j++ {
			xy := pixel.V(float64(i), float64(j)).Scaled(d.Pad).Add(d.Offset)
			d.Imd.Push(xy, xy.Add(xypad))
		}
		d.Imd.Rectangle(1)
	}
}
