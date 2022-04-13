package mandelbrot

// #include "mandelbrot.h"
import "C"

import (
	"image"
	"image/color"
	"runtime"
	"sync"
)

const (
	shiftStep                 = 30
	zoomCoefficient           = 1.2
	zoomIterationsCoefficient = 1.03
)

var orangeColor = (*[0]byte)(C.mandelbrot_color_orange)

type Mandelbrot struct {
	Width, Height  int
	Iterations     int
	Zoom           float64
	ShiftX, ShiftY float64
	Img            *image.RGBA
}

func New(width, height int) *Mandelbrot {
	m := &Mandelbrot{
		Width:      width,
		Height:     height,
		Iterations: 100,
		Zoom:       1.0,
		Img:        image.NewRGBA(image.Rect(0, 0, width, height)),
	}

	return m
}

func (m *Mandelbrot) CalculatePixelColor(x, y int) color.RGBA {
	// rgba_t mandelbrot_calculate_pixel_color(
	//		int32_t x,
	//		int32_t y,
	//		double shift_x,
	//		double shift_y,
	//		double zoomCoefficient,
	//		int32_t max_iterations,
	//		int32_t width,
	//		int32_t height,
	//		rgba_t(*color)(int32_t, int32_t)
	//	)
	var rgba C.rgba_t = C.mandelbrot_calculate_pixel_color(
		C.int(x),            // int32_t x
		C.int(y),            // int32_t y
		C.double(m.ShiftX),  // double shift_x
		C.double(m.ShiftY),  // double shift_y
		C.double(m.Zoom),    // double zoomCoefficient
		C.int(m.Iterations), // max_iterations
		C.int(m.Width),      // int32_t width
		C.int(m.Height),     // int32_t height
		orangeColor,         // rgba_t(*color)(int32_t, int32_t)
	)

	return color.RGBA{
		R: uint8(rgba.r),
		G: uint8(rgba.g),
		B: uint8(rgba.b),
		A: uint8(rgba.a),
	}
}

func (m *Mandelbrot) CalculateImage() *image.RGBA {
	numCPU := runtime.NumCPU()
	if numCPU < 2 {
		m.calculateImageSingleThread()
		return m.Img
	}

	m.calculateImageMultiThread(numCPU)
	return m.Img
}

func (m *Mandelbrot) calculateImageSingleThread() {
	for y := 0; y < m.Height; y++ {
		for x := 0; x < m.Width; x++ {
			m.Img.SetRGBA(x, y, m.CalculatePixelColor(x, y))
		}
	}
}

func (m *Mandelbrot) calculateImageMultiThread(goroutines int) {
	wg := sync.WaitGroup{}
	for i := 0; i < goroutines; i++ {
		wg.Add(1)
		from, to := m.Height/goroutines*i, m.Height/goroutines*(i+1)
		go func(from, to int) {
			for y := from; y < to; y++ {
				for x := 0; x < m.Width; x++ {
					m.Img.SetRGBA(x, y, m.CalculatePixelColor(x, y))
				}
			}

			wg.Done()
		}(from, to)
	}

	wg.Wait()
}

func (m *Mandelbrot) Image() *image.RGBA {
	return m.Img
}

func (m *Mandelbrot) ZoomIn() {
	m.Zoom *= zoomCoefficient
	m.ShiftX *= zoomCoefficient
	m.ShiftY *= zoomCoefficient

	m.Iterations = int(float64(m.Iterations) * zoomIterationsCoefficient)
}

func (m *Mandelbrot) ZoomOut() {
	m.Zoom /= zoomCoefficient
	m.ShiftX *= zoomCoefficient
	m.ShiftY *= zoomCoefficient

	m.Iterations = int(float64(m.Iterations) / zoomIterationsCoefficient)
}

func (m *Mandelbrot) ShiftUp() {
	m.ShiftY -= shiftStep
}

func (m *Mandelbrot) ShiftDown() {
	m.ShiftY += shiftStep
}

func (m *Mandelbrot) ShiftLeft() {
	m.ShiftX -= shiftStep
}

func (m *Mandelbrot) ShiftRight() {
	m.ShiftX += shiftStep
}
