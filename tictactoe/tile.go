package tictactoe

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
)

type tileSymbol string

const (
	tileEmpty  tileSymbol = ""
	tileCircle tileSymbol = "o"
	tileCross  tileSymbol = "x"

	tileSize   = 80
	tileMargin = 4
)

type Tile struct {
	image  *ebiten.Image
	symbol tileSymbol
}

func NewTile(symbol tileSymbol) *Tile {
	return &Tile{
		symbol: symbol,
		image:  ebiten.NewImage(tileSize, tileSize),
	}
}

func (t *Tile) Draw() {
	t.image.Fill(t.color())
}

func (t *Tile) Mark(symbol tileSymbol) {
	t.symbol = symbol
}

func (t *Tile) IsEmpty() bool {
	return t.symbol == tileEmpty
}

func (t *Tile) IsCross() bool {
	return t.symbol == tileCross
}

func (t *Tile) IsCircle() bool {
	return t.symbol == tileCircle
}

func (t *Tile) color() color.Color {
	switch {
	case t.IsEmpty():
		return tileEmptyColor
	case t.IsCircle():
		return tileCircleColor
	case t.IsCross():
		return tileCrossColor
	}
	panic("not reach")
}
