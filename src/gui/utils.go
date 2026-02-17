package gui

import (
	"image/color"
)

var colorMap = map[rune]color.NRGBA{
	'A': {255, 0, 0, 255},     // Red
	'B': {0, 255, 0, 255},     // Green
	'C': {0, 0, 255, 255},     // Blue
	'D': {255, 255, 0, 255},   // Yellow
	'E': {255, 0, 255, 255},   // Magenta
	'F': {0, 255, 255, 255},   // Cyan
	'G': {127, 0, 0, 255},     // Dark Red
	'H': {0, 127, 0, 255},     // Dark Green
	'I': {0, 0, 127, 255},     // Dark Blue
	'J': {127, 127, 0, 255},   // Olive
	'K': {127, 0, 127, 255},   // Purple
	'L': {0, 127, 127, 255},   // Teal
	'M': {191, 191, 191, 255}, // Silver
	'N': {255, 165, 0, 255},   // Orange
	'O': {165, 41, 41, 255},   // Brown
	'P': {138, 43, 227, 255},  // Blue Violet
	'Q': {94, 158, 161, 255},  // Cadet Blue
	'R': {209, 105, 31, 255},  // Chocolate
	'S': {255, 127, 79, 255},  // Coral
	'T': {99, 148, 237, 255},  // Cornflower Blue
	'U': {219, 20, 61, 255},   // Crimson
	'V': {0, 206, 209, 255},   // Dark Turquoise
	'W': {148, 0, 211, 255},   // Dark Violet
	'X': {255, 20, 147, 255},  // Deep Pink
	'Y': {33, 140, 33, 255},   // Forest Green
	'Z': {255, 215, 0, 255},   // Gold
}

func GetColor(char rune) color.Color {
	if c, ok := colorMap[char]; ok {
		return c
	}
	return color.NRGBA{R: 0, G: 0, B: 0, A: 255}
}
