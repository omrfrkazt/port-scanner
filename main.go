package main

import (
	"portScanner/cmd"

	"github.com/pterm/pterm"
)

func main() {

	pterm.DefaultBigText.WithLetters(
		pterm.NewLettersFromStringWithRGB("GO", pterm.NewRGB(0, 125, 156)),
		pterm.NewLettersFromStringWithRGB("SCAN", pterm.NewRGB(253, 221, 0))).
		Render()

	cmd.Execute()
}
