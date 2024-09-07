package tictactoe

import "image/color"

type symbol string

var (
	screenColor = color.RGBA{250, 248, 239, 255}
	frameColor  = color.RGBA{187, 173, 160, 255}

	symbolEmptyColor  = color.NRGBA{238, 228, 218, 89}
	symbolCircleColor = color.RGBA{243, 109, 114, 255}
	symbolCrossColor  = color.RGBA{90, 154, 243, 255}
)

func symbolColor(symbol symbol) color.Color {
	switch symbol {
	case symbolEmpty:
		return symbolEmptyColor
	case symbolCircle:
		return symbolCircleColor
	case symbolCross:
		return symbolCrossColor
	}
	panic("not reach")
}
