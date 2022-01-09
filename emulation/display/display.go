package display

import "image"

type Display interface {
	Run()
}

type Drawer interface {
	Draw(image.Image)
}
