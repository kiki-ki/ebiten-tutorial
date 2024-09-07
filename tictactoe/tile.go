package tictactoe

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
)

const (
	tileEmpty  = ""
	tileCircle = "o"
	tileCross  = "x"

	tileWidth  = 80
	tileHeight = 80
	tileMargin = 4
)

type Tile struct {
	image *ebiten.Image
	value string
}

func NewTile(value string) *Tile {
	return &Tile{
		value: value,
		image: ebiten.NewImage(tileWidth, tileHeight),
	}
}

func (t *Tile) isEmpty() bool {
	return t.value == tileEmpty
}

func (t *Tile) isCross() bool {
	return t.value == tileCross
}

func (t *Tile) isCircle() bool {
	return t.value == tileCircle
}

func (t *Tile) color() color.Color {
	switch {
	case t.isEmpty():
		return color.NRGBA{0xee, 0xe4, 0xda, 0x59}
	case t.isCircle():
		return color.RGBA{0xf9, 0xf6, 0xf2, 0xff}
	case t.isCross():
		return color.RGBA{0xff, 0xf6, 0xf2, 0xff}
	}
	panic("not reach")
}

func (t *Tile) Draw() {
	t.image.Fill(t.color())
}
