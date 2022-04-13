package main

import (
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
	"mandelbrot/mandelbrot"
)

const (
	WindowWidth  = 960
	WindowHeight = 600
)

func main() {
	pixelgl.Run(run)
}

func run() {
	cfg := pixelgl.WindowConfig{
		Title:  "Mandelbrot Set",
		Bounds: pixel.R(0, 0, WindowWidth, WindowHeight),
	}

	win, err := pixelgl.NewWindow(cfg)
	if err != nil {
		panic(err)
	}

	set := mandelbrot.New(WindowWidth, WindowHeight)
	for !win.Closed() {
		switch {
		case win.Pressed(pixelgl.KeyW):
			set.ShiftUp()
		case win.Pressed(pixelgl.KeyS):
			set.ShiftDown()
		case win.Pressed(pixelgl.KeyA):
			set.ShiftLeft()
		case win.Pressed(pixelgl.KeyD):
			set.ShiftRight()
		case win.Pressed(pixelgl.KeyLeftShift):
			set.ZoomIn()
		case win.Pressed(pixelgl.KeyLeftControl):
			set.ZoomOut()
		}

		pic := pixel.PictureDataFromImage(set.CalculateImage())
		pixel.NewSprite(pic, pic.Bounds()).Draw(win, pixel.IM.Moved(win.Bounds().Center()))

		win.Update()
	}
}
