package assets

import (
	"bytes"
	"embed"
	"github.com/hajimehoshi/ebiten/v2"
	"golang.org/x/image/font"
	"golang.org/x/image/font/opentype"
	"image"
)

var (

	//go:embed fonts/kenney-high-square.ttf
	normalFont []byte

	//go:embed *
	assetsFS embed.FS

	SmallFont  = mustLoadFont(normalFont, 18)
	NormalFont = mustLoadFont(normalFont, 32)
)

func mustLoadFont(data []byte, size int) font.Face {
	f, err := opentype.Parse(data)
	if err != nil {
		panic(err)
	}

	face, err := opentype.NewFace(f, &opentype.FaceOptions{
		Size:    float64(size),
		DPI:     72,
		Hinting: font.HintingFull,
	})
	if err != nil {
		panic(err)
	}

	return face
}

func mustNewEbitenImage(data []byte) *ebiten.Image {
	img, _, err := image.Decode(bytes.NewReader(data))
	if err != nil {
		panic(err)
	}

	return ebiten.NewImageFromImage(img)
}
