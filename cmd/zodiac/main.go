package main

import (
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/widget"
	"github.com/Soulsbane/internal/zodiac/zodiac"
)

func main() {

	for _, yr := range []int{1935, 1938, 1968, 1972, 1976} {
        a, yy, e, sb, cy := zodiac.GetSign(yr)
        fmt.Printf("%d: %s %s, %s, Cycle year %d %s\n",
            yr, e, a, yy, cy, sb)
    }
	a := app.New()
	w := a.NewWindow("Hello World")

	w.SetContent(widget.NewLabel("Hello World!"))
	w.ShowAndRun()
}
