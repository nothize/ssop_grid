package main

import (
	"github.com/faiface/pixel/pixelgl"
	"io/ioutil"
	"log"
)

func main() {
	cfg, err := ioutil.ReadFile("gridconfig.yaml")
	if err != nil {
		log.Fatal("Config file not found")
	}
	d := NewDraw(NewGridConfig(string(cfg)))

	pixelgl.Run(d.run)
}