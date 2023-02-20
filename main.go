package main

import (
	"fmt"

	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
	"github.com/nazarslota/mandelbrot/fractal"
)

const (
	mandelbrotWidth  = 1280
	mandelbrotHeight = 720

	mandelbrotZoomCoefficient = 2
)

var mandelbrot = fractal.NewMandelbrot(mandelbrotWidth, mandelbrotHeight, 1024)

func main() {
	a := app.New()
	a.Settings().SetTheme(theme.DefaultTheme())

	w := a.NewWindow("Mandelbrot Set")
	w.SetFixedSize(true)

	mandelbrotImage := canvas.NewImageFromImage(mandelbrot.BuildImage())
	mandelbrotImage.FillMode = canvas.ImageFillOriginal

	zoomLabel := widget.NewLabel(fmt.Sprintf("Zoom: %.4f", mandelbrot.GetZoom()))
	zoomInButton := widget.NewButtonWithIcon("Zoom In", theme.ZoomInIcon(), func() {
		mandelbrot.SetZoom(mandelbrot.GetZoom() * mandelbrotZoomCoefficient)
		mandelbrot.Refresh()

		zoomLabel.SetText(fmt.Sprintf("Zoom: %.4f", mandelbrot.GetZoom()))
		zoomLabel.Refresh()

		mandelbrotImage.Image = mandelbrot.BuildImage()
		mandelbrotImage.Refresh()
	})
	zoomOutButton := widget.NewButtonWithIcon("Zoom Out", theme.ZoomOutIcon(), func() {
		mandelbrot.SetZoom(mandelbrot.GetZoom() / mandelbrotZoomCoefficient)
		mandelbrot.Refresh()

		zoomLabel.SetText(fmt.Sprintf("Zoom: %.4f", mandelbrot.GetZoom()))
		zoomLabel.Refresh()

		mandelbrotImage.Image = mandelbrot.BuildImage()
		mandelbrotImage.Refresh()
	})

	shiftXLabel := widget.NewLabel(fmt.Sprintf("X: %.4f", mandelbrot.GetShiftX()))
	moveRightButton := widget.NewButtonWithIcon("Move Right", theme.NavigateNextIcon(), func() {
		mandelbrot.SetShiftX(mandelbrot.GetShiftX() + 1.0/mandelbrot.GetZoom())
		mandelbrot.Refresh()

		shiftXLabel.SetText(fmt.Sprintf("X: %.4f", mandelbrot.GetShiftX()))
		shiftXLabel.Refresh()

		mandelbrotImage.Image = mandelbrot.BuildImage()
		mandelbrotImage.Refresh()
	})
	moveLeftButton := widget.NewButtonWithIcon("Move Left", theme.NavigateBackIcon(), func() {
		mandelbrot.SetShiftX(mandelbrot.GetShiftX() - 1.0/mandelbrot.GetZoom())
		mandelbrot.Refresh()

		shiftXLabel.SetText(fmt.Sprintf("X: %.4f", mandelbrot.GetShiftX()))
		shiftXLabel.Refresh()

		mandelbrotImage.Image = mandelbrot.BuildImage()
		mandelbrotImage.Refresh()
	})

	shiftYLabel := widget.NewLabel(fmt.Sprintf("Y: %.4f", mandelbrot.GetShiftY()))
	moveUpButton := widget.NewButtonWithIcon("Move Up", theme.MoveUpIcon(), func() {
		mandelbrot.SetShiftY(mandelbrot.GetShiftY() - 1.0/mandelbrot.GetZoom())
		mandelbrot.Refresh()

		shiftYLabel.SetText(fmt.Sprintf("Y: %.4f", mandelbrot.GetShiftY()))
		shiftYLabel.Refresh()

		mandelbrotImage.Image = mandelbrot.BuildImage()
		mandelbrotImage.Refresh()
	})
	moveDownButton := widget.NewButtonWithIcon("Move Down", theme.MoveDownIcon(), func() {
		mandelbrot.SetShiftY(mandelbrot.GetShiftY() + 1.0/mandelbrot.GetZoom())
		mandelbrot.Refresh()

		shiftYLabel.SetText(fmt.Sprintf("Y: %.4f", mandelbrot.GetShiftY()))
		shiftYLabel.Refresh()

		mandelbrotImage.Image = mandelbrot.BuildImage()
		mandelbrotImage.Refresh()
	})

	w.SetContent(container.NewHBox(
		container.NewCenter(mandelbrotImage),
		container.NewVBox(
			container.NewHBox(zoomInButton, zoomOutButton),
			container.NewHBox(moveUpButton, moveDownButton),
			container.NewHBox(moveLeftButton, moveRightButton),
			zoomLabel,
			container.NewHBox(shiftXLabel, shiftYLabel),
		),
	))

	w.Show()
	a.Run()
}
