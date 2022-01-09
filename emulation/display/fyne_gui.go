package display

import (
	"image"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
)

type FyneGUI struct {
	app    *fyne.App
	window *fyne.Window
	image  fyne.CanvasObject
}

func NewFyneGUI() *FyneGUI {
	app := app.New()
	window := app.NewWindow("Proto-GUI")

	return &FyneGUI{
		app:    &app,
		window: &window,
	}
}

func (fg *FyneGUI) Draw(frame image.Image) {
	size := fyne.NewSize(160, 144)

	fg.image = canvas.NewRasterFromImage(frame)
	fg.image.Resize(size)

	(*fg.window).SetContent(fg.image)
}

func (fg *FyneGUI) Run() {
	size := fyne.NewSize(640, 576)

	fg.image = canvas.NewImageFromImage(nil)

	(*fg.window).Resize(size)
	(*fg.window).SetContent(fg.image)
	(*fg.window).ShowAndRun()
}
