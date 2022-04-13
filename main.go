package main

import (
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
	"mandelbrot/mandelbrot"
)

// Pixel window width and height.
const (
	WindowWidth  = 960
	WindowHeight = 600
)

func main() {
	pixelgl.Run(run)
}

func run() {
	cfg := pixelgl.WindowConfig{
		Title:       "Mandelbrot Set",
		Bounds:      pixel.R(0, 0, WindowWidth, WindowHeight),
		AlwaysOnTop: true,
	}

	win, err := pixelgl.NewWindow(cfg)
	if err != nil {
		panic(err)
	}

	set := mandelbrot.New(WindowWidth, WindowHeight)
	for !win.Closed() {
		switch {
		case win.JustPressed(pixelgl.KeyW):
			set.ShiftUp()
		case win.JustPressed(pixelgl.KeyS):
			set.ShiftDown()
		case win.JustPressed(pixelgl.KeyA):
			set.ShiftLeft()
		case win.JustPressed(pixelgl.KeyD):
			set.ShiftRight()
		case win.JustPressed(pixelgl.KeyLeftShift):
			set.ZoomIn()
		case win.JustPressed(pixelgl.KeyLeftControl):
			set.ZoomOut()
		}

		pic := pixel.PictureDataFromImage(set.CalculateImage())
		pixel.NewSprite(pic, pic.Bounds()).Draw(win, pixel.IM.Moved(win.Bounds().Center()))

		win.Update()
	}
}
