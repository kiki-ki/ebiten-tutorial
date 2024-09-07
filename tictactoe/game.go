package tictactoe

import (
	"github.com/hajimehoshi/ebiten/v2"
)

const (
	ScreenWidth  = 420
	ScreenHeight = 600
)

type Game struct {
	currentPlayer *Player
	players       []*Player
	board         *Board
}

func NewGame() *Game {
	players := []*Player{
		NewPlayer(tileCircle),
		NewPlayer(tileCross),
	}
	return &Game{
		currentPlayer: players[0],
		players:       players,
		board:         NewBoard(),
	}
}

func (g *Game) Update() error {
	if ebiten.IsMouseButtonPressed(ebiten.MouseButtonLeft) {
		mouseX, mouseY := ebiten.CursorPosition()

		if g.board.IsPressed(mouseX, mouseY) {
			row := (mouseX - boardMinX) / tileSize
			column := (mouseY - boardMinY) / tileSize

			if !g.board.IsValidEmptyTile(row, column) {
				return nil
			}

			selected := g.board.tiles[row][column]
			selected.Mark(g.currentPlayer.symbol)

			if g.board.HasWinningLine() {
				g.processWin(g.currentPlayer)
			}
			if g.board.IsAllMarked() {
				g.processDraw()
			}

			g.changePlayer()
		}
	}

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

func (g *Game) processWin(winner *Player) {}

func (g *Game) processDraw() {}

func (g *Game) changePlayer() {
	for _, player := range g.players {
		if g.currentPlayer.symbol != player.symbol {
			g.currentPlayer = player
		}
	}
}
