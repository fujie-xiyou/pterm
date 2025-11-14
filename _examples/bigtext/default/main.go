package main

import (
	"github.com/fujie-xiyou/pterm"
	"github.com/fujie-xiyou/pterm/putils"
)

func main() {
	// Define the text to be rendered
	var text = "PTerm"

	// Convert the text into a format suitable for PTerm
	var letters = putils.LettersFromString(text)

	// Render the text using PTerm's default big text style
	pterm.DefaultBigText.WithLetters(letters).Render()
}
