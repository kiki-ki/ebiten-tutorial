package tictactoe

type Player struct {
	symbol tileSymbol
}

func NewPlayer(symbol tileSymbol) *Player {
	return &Player{
		symbol: symbol,
	}
}
