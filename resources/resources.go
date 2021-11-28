package resources

import (
	"embed"
	"image/color"

	"github.com/blizzy78/ebitenui/image"
	"github.com/blizzy78/ebitenui/widget"
	"github.com/golang/freetype/truetype"
	"golang.org/x/image/font"
)

var Font = loadFont("NotoSans-Regular", truetype.Options{Size: 12})

var SelectSandButtonImage = &widget.ButtonImage{
	Idle:     image.NewNineSliceColor(color.RGBA{R: 0, G: 0, B: 0, A: 255}),
	Hover:    image.NewNineSliceColor(color.RGBA{R: 0, G: 0, B: 0, A: 255}),
	Pressed:  image.NewNineSliceColor(color.RGBA{R: 0, G: 0, B: 0, A: 255}),
	Disabled: image.NewNineSliceColor(color.RGBA{R: 25, G: 25, B: 25, A: 255}),
}

var SelectSandButtonText = &widget.ButtonTextColor{
	Idle:     color.RGBA{R: 192, G: 192, B: 192, A: 255},
	Disabled: color.RGBA{R: 255, G: 255, B: 255, A: 255},
}

var HeaderTextFont = loadFont("NotoSans-Bold", truetype.Options{Size: 14})
var HeaderTextColor = &widget.LabelColor{
	Idle:     color.RGBA{R: 255, G: 255, B: 255, A: 255},
	Disabled: color.RGBA{R: 255, G: 255, B: 255, A: 255},
}

// ---------------------------------------------------------------------------------------------------------------------

// loadFont loads an embedded font.
func loadFont(name string, options truetype.Options) font.Face {
	font, err := truetype.Parse(loadFile("NotoSans-Regular.ttf"))
	if err != nil {
		panic("failed to decode font file: " + err.Error())
	}

	return truetype.NewFace(font, &options)
}

// loadFile loads an embedded file.
func loadFile(name string) []byte {
	contents, err := _files.ReadFile(name)
	if err != nil {
		panic("failed to open file: " + err.Error())
	}

	return contents
}

//go:embed *
var _files embed.FS
