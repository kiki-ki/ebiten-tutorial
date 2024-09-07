package tictactoe

import (
	"bytes"
	"log"

	"github.com/hajimehoshi/ebiten/examples/resources/fonts"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text/v2"
)

const (
	ScreenWidth  = 420
	ScreenHeight = 600
)

var (
	mplusFaceSource *text.GoTextFaceSource
)

func init() {
	// load font data
	s, err := text.NewGoTextFaceSource(bytes.NewReader(fonts.MPlus1pRegular_ttf))
	if err != nil {
		log.Fatal(err)
	}
	mplusFaceSource = s
}

type Game struct {
	current *Player
	players []*Player
	board   *Board
}

func NewGame() *Game {
	players := []*Player{
		NewPlayer(symbolCircle),
		NewPlayer(symbolCross),
	}
	return &Game{
		current: players[0],
		players: players,
		board:   NewBoard(),
	}
}

func (g *Game) Reset() {
	g.board = NewBoard()
	g.current = g.players[0]
}

func (g *Game) Update() error {
	if !g.isFinished() && ebiten.IsMouseButtonPressed(ebiten.MouseButtonLeft) {
		mouseX, mouseY := ebiten.CursorPosition()

		if g.board.IsPressed(mouseX, mouseY) {
			row := (mouseX - boardMinX) / tileSize
			column := (mouseY - boardMinY) / tileSize

			if !g.board.IsValidEmptyTile(row, column) {
				return nil
			}

			selected := g.board.tiles[row][column]
			selected.Mark(g.current.symbol)

			if !g.isFinished() {
				g.changePlayer()
			}
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

	if g.isFinished() {
		fontFace := &text.GoTextFace{
			Source: mplusFaceSource,
			Size:   24,
		}
		textOpts := &text.DrawOptions{}
		textOpts.PrimaryAlign = text.AlignCenter
		textOpts.GeoM.Translate(float64(ScreenWidth/2), float64(ScreenHeight/8))

		if g.winner() != nil {
			textOpts.ColorScale.ScaleWithColor(symbolColor(g.winner().symbol))
			text.Draw(screen, g.winner().WinnerMsg(), fontFace, textOpts)
		} else if g.board.IsAllMarked() {
			textOpts.ColorScale.ScaleWithColor(frameColor)
			text.Draw(screen, "draw!", fontFace, textOpts)
		}

		if ebiten.IsMouseButtonPressed(ebiten.MouseButtonRight) {
			g.Reset()
		}
	}
}

func (g *Game) isFinished() bool {
	return g.winner() != nil || g.board.IsAllMarked()
}

func (g *Game) winner() *Player {
	if g.board.HasWinningLine() {
		return g.current
	}
	return nil
}

func (g *Game) changePlayer() {
	for _, player := range g.players {
		if g.current.symbol != player.symbol {
			g.current = player
			return
		}
	}
}
