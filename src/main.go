package main

import (
	"log"

	"github.com/rivo/tview"
)

func main() {
	app := tview.NewApplication()
	pages := tview.NewPages()

	// initialization of pages
	pagesInitHome(pages)

	// starting
	err := app.SetRoot(pages, true).EnableMouse(true).Run()
	if err != nil {
		log.Println(err)
	}
}
