package resources

import (
	"embed"

	"github.com/golang/freetype/truetype"
	"golang.org/x/image/font"
)

var Font font.Face

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
