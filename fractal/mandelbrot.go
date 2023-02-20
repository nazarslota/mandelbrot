package fractal

// #include "mandelbrot.h"
import "C"

import (
	"image"
	"image/color"
	"runtime"
	"sync"
)

type Mandelbrot struct {
	width, height int
	maxiterations int

	pointiterations []int
	img             *image.RGBA

	zoom           float64
	shiftx, shifty float64

	pixelwidth, pixelheight float64
}

const (
	MandelbrotDefaultMaxIterations = 256
	MandelbrotDefaultZoom          = 1

	xmax, xmin = 2.0, -2.0
	ymax, ymin = 2.0, -2.0
)

func NewMandelbrot(width, height int, maxiterations ...int) *Mandelbrot {
	if width <= 0 {
		panic("width must be greater than zero")
	} else if height <= 0 {
		panic("height must be greater than zero")
	}

	s := &Mandelbrot{
		width:           width,
		height:          height,
		maxiterations:   MandelbrotDefaultMaxIterations,
		pointiterations: make([]int, width*height),
		img:             image.NewRGBA(image.Rect(0, 0, width, height)),
		zoom:            MandelbrotDefaultZoom,
		pixelwidth:      (xmax - xmin) / float64(width),
		pixelheight:     (ymax - ymin) / float64(height),
	}

	if len(maxiterations) > 0 {
		if maxiterations[0] <= 0 {
			panic("maxiterations must be greater than zero")
		}
		s.maxiterations = maxiterations[0]
	}
	s.Refresh()
	return s
}

func (s *Mandelbrot) Refresh() {
	numCPU := runtime.NumCPU()
	if numCPU < 2 {
		s.refresh(1)
	}
	s.refresh(numCPU)
}

func (s *Mandelbrot) refresh(goroutines int) {
	if goroutines <= 0 {
		panic("goroutines count must be greater than zero")
	}

	var wg sync.WaitGroup
	wg.Add(goroutines)

	for i := 0; i < goroutines; i++ {
		from := s.width / goroutines * i
		to := s.width / goroutines * (i + 1)

		go func(from, to int) {
			for y := 0; y < s.height; y++ {
				for x := from; x < to; x++ {
					pi := y*s.width + x
					s.pointiterations[pi] = s.iterations(x, y)
				}
			}
			wg.Done()
		}(from, to)
	}
	wg.Wait()
}

func (s *Mandelbrot) iterations(x, y int) int {
	nx := xmin + float64(x)*s.pixelwidth
	ny := ymin + float64(y)*s.pixelheight

	var iterations C.int = C.mandelbrot_iterations(
		C.int(s.width),         // int w
		C.int(s.height),        // int h
		C.double(nx),           // double x
		C.double(ny),           // double y
		C.double(s.shiftx),     // double shift_x
		C.double(s.shifty),     // double shift_y
		C.double(s.zoom),       // double zoom
		C.int(s.maxiterations), // int max_iterations
	)
	return int(iterations)
}

func (s *Mandelbrot) BuildImage() image.Image {
	numCPU := runtime.NumCPU()
	if numCPU < 2 {
		s.buildimage(1)
	}
	return s.buildimage(numCPU)
}

func (s *Mandelbrot) buildimage(goroutines int) image.Image {
	var wg sync.WaitGroup
	wg.Add(goroutines)

	np := s.width * s.height
	for i := 0; i < goroutines; i++ {
		from := np / goroutines * i
		to := np / goroutines * (i + 1)

		go func(from, to int) {
			for i := from; i < to; i++ {
				x := i % s.width
				y := i / s.width

				c := s.color(s.pointiterations[i])
				s.img.SetRGBA(x, y, c)
			}
			wg.Done()
		}(from, to)
	}
	wg.Wait()
	return s.img
}

func (s *Mandelbrot) color(iterations int) color.RGBA {
	var c C.rgba_t = C.mandelbrot_color(
		C.int(iterations),      // int iterations
		C.int(s.maxiterations), // int max_iterations
	)
	return color.RGBA{R: uint8(c.r), G: uint8(c.g), B: uint8(c.b), A: uint8(c.a)}
}

func (s *Mandelbrot) GetMaxIterations() int           { return s.maxiterations }
func (s *Mandelbrot) SetMaxIterations(iterations int) { s.maxiterations = iterations }

func (s *Mandelbrot) GetZoom() float64     { return s.zoom }
func (s *Mandelbrot) SetZoom(zoom float64) { s.zoom = zoom }

func (s *Mandelbrot) GetShiftX() float64      { return s.shiftx }
func (s *Mandelbrot) SetShiftX(shift float64) { s.shiftx = shift }

func (s *Mandelbrot) GetShiftY() float64      { return s.shifty }
func (s *Mandelbrot) SetShiftY(shift float64) { s.shifty = shift }
