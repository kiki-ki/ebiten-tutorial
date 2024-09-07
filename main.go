package main

import (
	"ebiten-tutorial/tictactoe"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
)

func main() {
	game := tictactoe.NewGame()

	ebiten.SetWindowSize(tictactoe.ScreenWidth, tictactoe.ScreenHeight)
	ebiten.SetWindowTitle("TicTacToe")

	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}
}
