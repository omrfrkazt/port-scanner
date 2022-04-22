package main

import (
	"fmt"
	"portScanner/port"
	"time"

	"github.com/pterm/pterm"
)

func main() {

	pterm.DefaultBigText.WithLetters(
		pterm.NewLettersFromStringWithRGB("GO", pterm.NewRGB(0, 125, 156)),
		pterm.NewLettersFromStringWithRGB("SCAN", pterm.NewRGB(253, 221, 0))).
		Render()

	channel := make(chan []port.ScanResult)
	fmt.Println("Started: ", time.Now().Local().UTC())
	go port.InitialScan("google.com", 1, 1024, &channel)

	fmt.Println(<-channel)
	fmt.Println("Finished: ", time.Now().Local().UTC())
}
