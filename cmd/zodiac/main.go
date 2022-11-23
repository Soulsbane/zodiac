package main

import (
	"fmt"

	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/widget"
	"github.com/Soulsbane/zodiac/internal/zodiac"
)

func main() {

	for _, yr := range []int{1935, 1938, 1968, 1972, 1976} {
		sign := zodiac.GetSign(yr)
		fmt.Printf("%d: %s %s, %s, Cycle year %d %s\n",
			sign.Year, sign.Animal, sign.YinYang, sign.Element, sign.Year, sign.StemBranch)
	}
	a := app.New()
	w := a.NewWindow("Hello World")

	w.SetContent(widget.NewLabel("Hello World!"))
	w.ShowAndRun()
}
