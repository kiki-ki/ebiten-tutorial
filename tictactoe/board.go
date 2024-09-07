package tictactoe

import (
	"github.com/hajimehoshi/ebiten/v2"
)

const (
	boardLength = 3

	boardWidth = boardLength*tileSize + (boardLength+1)*tileMargin
	boardMinX  = (ScreenWidth - boardWidth) / 2
	boardMaxX  = boardMinX + boardWidth

	boardHeight = boardLength*tileSize + (boardLength+1)*tileMargin
	boardMinY   = (ScreenHeight - boardHeight) / 2
	boardMaxY   = boardMinY + boardHeight
)

type Board struct {
	image *ebiten.Image
	tiles [][]*Tile
}

func NewBoard() *Board {
	tiles := make([][]*Tile, boardLength)
	for i := range tiles {
		tiles[i] = make([]*Tile, boardLength)
		for j := range tiles[i] {
			tiles[i][j] = NewTile(tileEmpty)
		}
	}
	board := &Board{
		tiles: tiles,
	}
	board.image = ebiten.NewImage(boardWidth, boardHeight)
	return board
}

func (b *Board) Draw() {
	b.image.Fill(frameColor)

	for i := 0; i < boardLength; i++ {
		for j := 0; j < boardLength; j++ {
			tile := b.tiles[i][j]
			tile.Draw()

			x := i*tileSize + (i+1)*tileMargin
			y := j*tileSize + (j+1)*tileMargin
			opts := &ebiten.DrawImageOptions{}
			opts.GeoM.Translate(float64(x), float64(y))

			b.image.DrawImage(tile.image, opts)
		}
	}
}

func (b *Board) HasWinningLine() bool {
	for i := 0; i < boardLength; i++ {
		// Check row
		row := b.tiles[i]
		if b.isWinningLine(row) {
			return true
		}
		// Check column
		column := make([]*Tile, boardLength)
		for j := 0; j < boardLength; j++ {
			column[j] = b.tiles[j][i]
		}
		if b.isWinningLine(column) {
			return true
		}
	}

	// Check diagonal (\)
	diag1 := make([]*Tile, boardLength)
	for i := 0; i < boardLength; i++ {
		diag1[i] = b.tiles[i][i]
	}
	if b.isWinningLine(diag1) {
		return true
	}

	// Check diagonal (/)
	diag2 := make([]*Tile, boardLength)
	for i := 0; i < boardLength; i++ {
		diag2[i] = b.tiles[i][boardLength-i-1]
	}

	return b.isWinningLine(diag2)
}

func (b *Board) isWinningLine(row []*Tile) bool {
	first := row[0]
	if first.IsEmpty() {
		return false
	}
	for _, tile := range row {
		if tile.symbol != first.symbol {
			return false
		}
	}
	return true
}

func (b *Board) IsAllMarked() bool {
	for _, row := range b.tiles {
		for _, tile := range row {
			if tile.IsEmpty() {
				return false
			}
		}
	}
	return true
}

func (b *Board) IsPressed(mouseX, mouseY int) bool {
	isPressedX := mouseX >= boardMinX && mouseX < boardMaxX
	isPressedY := mouseY >= boardMinY && mouseY < boardMaxY
	return isPressedX && isPressedY
}

func (b *Board) IsValidEmptyTile(row, column int) bool {
	isValidRow := row >= 0 && row < boardLength
	isValidColumn := column >= 0 && column < boardLength
	return isValidRow && isValidColumn && b.tiles[row][column].IsEmpty()
}
