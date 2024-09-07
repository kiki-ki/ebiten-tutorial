package tictactoe

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
)

var (
	frameColor = color.RGBA{0xbb, 0xad, 0xa0, 0xff}
)

type Board struct {
	image  *ebiten.Image
	length int
	tiles  [][]*Tile
}

func NewBoard(length int) *Board {
	row := make([]*Tile, length)
	for i := range row {
		row[i] = NewTile(tileEmpty)
	}
	tiles := make([][]*Tile, length)
	for i := range tiles {
		tiles[i] = row
	}
	board := &Board{
		length: length,
		tiles:  tiles,
	}
	board.image = ebiten.NewImage(board.size())
	return board
}

func (b *Board) size() (int, int) {
	x := b.length*tileWidth + (b.length+1)*tileMargin
	y := b.length*tileHeight + (b.length+1)*tileMargin
	return x, y
}

func (b *Board) Draw() {
	b.image.Fill(frameColor)

	for i := 0; i < b.length; i++ {
		for j := 0; j < b.length; j++ {
			tile := b.tiles[i][j]
			tile.Draw()

			x := i*tileWidth + (i+1)*tileMargin
			y := j*tileHeight + (j+1)*tileMargin
			opts := &ebiten.DrawImageOptions{}
			opts.GeoM.Translate(float64(x), float64(y))

			b.image.DrawImage(tile.image, opts)
		}
	}
}
