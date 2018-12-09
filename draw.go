package main

import (
	"github.com/faiface/pixel/pixelgl"
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/imdraw"
	"golang.org/x/image/colornames"
	"time"
	"github.com/sirupsen/logrus"
)

type D struct {
	GridConfig *GridConfig
	Size       int
	Cfg        *pixelgl.WindowConfig
	Imd        *imdraw.IMDraw
	Win        *pixelgl.Window
	Pad        float64
	Offset     pixel.Vec
}

func (d *D) run() {
	d.Init()

	imd := d.Imd

	d.DrawBoard()
	t := time.NewTicker(time.Second / 30)
	for !d.Win.Closed() {
		imd.Draw(d.Win)
		d.Win.Update()
		<- t.C
	}
}

func NewDraw(gc *GridConfig) *D {
	d := &D{GridConfig: gc, Size: gc.Size}
	return d
}

func (d *D) Init() {
	var err error
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
	xy := pixel.V(float64(0), float64(0)).Scaled(d.Pad).Add(d.Offset)
	xy2 := pixel.V(float64(d.Size), float64(d.Size)).Scaled(d.Pad).Add(d.Offset)
	d.Imd.Push(xy, xy2)
	d.Imd.Rectangle(4)

	h := d.GridConfig.GetHorizontal()
	for i := 0; i < len(h); i++ {
		for j := 0; j < len(h[i]); j++ {
			if h[i][j] {
				x := j
				y := d.GridConfig.Size - (i + 1)
				xy := pixel.V(float64(x), float64(y)).Scaled(d.Pad).Add(d.Offset)
				d.Imd.Push(xy, xy.Add(pixel.Vec{ d.Pad, 0}))
				d.Imd.Line(4)
			}
		}
	}
	v := d.GridConfig.GetVerticle()
	for i := 0; i < len(v); i++ {
		for j := 0; j < len(v[i]); j++ {
			if v[i][j] {
				logrus.Infof("%d, %d = %b", i, j, v[i][j])
				x := i + 1
				y := d.GridConfig.Size - (j + 1)
				xy := pixel.V(float64(x), float64(y)).Scaled(d.Pad).Add(d.Offset)
				d.Imd.Push(xy, xy.Add(pixel.Vec{ 0, d.Pad}))
				d.Imd.Line(4)
			}
		}
	}
}
