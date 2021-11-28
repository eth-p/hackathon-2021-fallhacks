package resources

import (
	"embed"
	"image/color"

	"github.com/blizzy78/ebitenui/image"
	"github.com/blizzy78/ebitenui/widget"
	"github.com/golang/freetype/truetype"
	"golang.org/x/image/font"
)

var Font font.Face

var SelectSandCheckbox = &widget.CheckboxGraphicImage{
	Unchecked: &widget.ButtonImageImage{
		Idle:     image.NewImageColor(color.RGBA{R: 0, G: 0, B: 0, A: 255}),
		Disabled: image.NewImageColor(color.RGBA{R: 0, G: 0, B: 0, A: 255}),
	},
	Checked: &widget.ButtonImageImage{
		Idle:     image.NewImageColor(color.RGBA{R: 0, G: 0, B: 0, A: 255}),
		Disabled: image.NewImageColor(color.RGBA{R: 0, G: 0, B: 0, A: 255}),
	},
	Greyed: &widget.ButtonImageImage{
		Idle:     image.NewImageColor(color.RGBA{R: 0, G: 0, B: 0, A: 255}),
		Disabled: image.NewImageColor(color.RGBA{R: 0, G: 0, B: 0, A: 255}),
	},
}

var SelectSandCheckboxBtn = &widget.ButtonImage{
	Idle:     image.NewNineSliceColor(color.RGBA{R: 0, G: 0, B: 0, A: 255}),
	Hover:    image.NewNineSliceColor(color.RGBA{R: 0, G: 0, B: 0, A: 255}),
	Pressed:  image.NewNineSliceColor(color.RGBA{R: 0, G: 0, B: 0, A: 255}),
	Disabled: image.NewNineSliceColor(color.RGBA{R: 0, G: 0, B: 0, A: 255}),
}

var SelectSandLabel = &widget.LabelColor{
	Idle:     color.RGBA{R: 255, A: 255},
	Disabled: color.RGBA{R: 255, A: 255},
}

func init() {
	font, _ := truetype.Parse(file("NotoSans-Regular.ttf"))
	Font = truetype.NewFace(font, &truetype.Options{
		Size: 12,
	})
}

func file(name string) []byte {
	contents, err := _files.ReadFile(name)
	if err != nil {
		panic("failed to open file: " + err.Error())
	}

	return contents
}

//go:embed *
var _files embed.FS

type Config struct {
	UpdateInterval uint
	Paused         bool
}
