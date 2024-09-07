package tictactoe

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
)

const (
	symbolEmpty  symbol = ""
	symbolCircle symbol = "o"
	symbolCross  symbol = "x"

	tileSize   = 80
	tileMargin = 4
)

type Tile struct {
	image  *ebiten.Image
	symbol symbol
}

func NewTile(symbol symbol) *Tile {
	return &Tile{
		symbol: symbol,
		image:  ebiten.NewImage(tileSize, tileSize),
	}
}

func (t *Tile) Draw() {
	t.image.Fill(t.color())
}

func (t *Tile) Mark(symbol symbol) {
	t.symbol = symbol
}

func (t *Tile) IsEmpty() bool {
	return t.symbol == symbolEmpty
}

func (t *Tile) IsCross() bool {
	return t.symbol == symbolCross
}

func (t *Tile) IsCircle() bool {
	return t.symbol == symbolCircle
}

func (t *Tile) color() color.Color {
	return symbolColor(t.symbol)
}
