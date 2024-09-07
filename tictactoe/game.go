package tictactoe

import (
	"github.com/hajimehoshi/ebiten/v2"
)

const (
	ScreenWidth  = 420
	ScreenHeight = 600

	boardLength = 3
)

type Game struct {
	board *Board
}

func NewGame() *Game {
	board := NewBoard(boardLength)
	return &Game{
		board: board,
	}
}

func (g *Game) Update() error {
	return nil
}

func (g *Game) Layout(width, height int) (screenWidth, screenHeight int) {
	return ScreenWidth, ScreenHeight
}

func (g *Game) Draw(screen *ebiten.Image) {
	screen.Fill(screenColor)
	g.board.Draw()

	opts := &ebiten.DrawImageOptions{}
	sw, sh := screen.Bounds().Dx(), screen.Bounds().Dy()
	bw, bh := g.board.image.Bounds().Dx(), g.board.image.Bounds().Dy()
	x := (sw - bw) / 2
	y := (sh - bh) / 2
	opts.GeoM.Translate(float64(x), float64(y))

	screen.DrawImage(g.board.image, opts)
}
