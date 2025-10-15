package main

import (
	"fmt"
	"syscall/js"

	"github.com/J-Nokwal/gocade/pkg/utils"
)

func main() {
	// Prevent the program from exiting
	done := make(chan struct{}, 0)
	// // Register functions to be called from JavaScript
	// js.Global().Set("goInit", js.FuncOf(initEngine))
	initEngine(
		GameWindow{
			// AspectRatio: utils.PtrFloat64(16.0 / 9.0),
			// FillMaxSize: true,
		},
	)

	<-done
}

type GameWindow struct {
	// AspectRatio *float64
	Width  *int
	Height *int
	// FillMaxSize bool
	Canvas js.Value
}

func initEngine(gameWindow GameWindow) {
	// Initialize your engine here
	fmt.Println("Engine initialized")
	// Get Size of body
	body := js.Global().Get("document").Get("body")
	body.Set("style", "width: 100dvw; height: 100dvh; margin: 0; padding: 0; overflow: hidden; background-color: blue;")
	gameWindow.Width = utils.PtrInt(body.Get("clientWidth").Int())
	gameWindow.Height = utils.PtrInt(body.Get("clientHeight").Int())

	gameWindow.Canvas = js.Global().Get("document").Call("createElement", "canvas")
	gameWindow.Canvas.Set("id", "gameCanvas")
	gameWindow.Canvas.Set("width", *gameWindow.Width)
	gameWindow.Canvas.Set("height", *gameWindow.Height)
	// add gameWindow to body
	body.Call("appendChild", gameWindow.Canvas)
}
