package tictactoe

import (
	"fmt"
)

type Player struct {
	symbol symbol
}

func NewPlayer(symbol symbol) *Player {
	return &Player{
		symbol: symbol,
	}
}

func (p *Player) DrawWin() {

}

func (p *Player) WinnerMsg() string {
	return fmt.Sprintf("%s wins!", p.symbol)
}
