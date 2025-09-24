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
			AspectRatio: utils.PtrFloat64(16.0 / 9.0),
			FillMaxSize: true,
		},
	)
	<-done
}

// func initEngine(this js.Value, args []js.Value) interface{} {
// 	// Initialize your engine here

// 	return nil
// }

type GameWindow struct {
	AspectRatio *float64
	Width       *int
	Height      *int
	FillMaxSize bool
	Canvas      js.Value
}

func initEngine(gameWindow GameWindow) {
	// Initialize your engine here
	fmt.Println("Engine initialized")

	// Create a WebGL canvas and append it to the document
	gameWindow.Canvas = js.Global().Get("document").Call("createElement", "canvas")
	gameWindow.Canvas.Set("id", "gameCanvas2")

	body := js.Global().Get("document").Get("body")

	body.Set("style", map[string]interface{}{
		"border":     "1px solid black",
		"background": "pink",
		"width":      "100dvh",
		"height":     "100dvh",
	})
	body.Get("gameCanvas", gameWindow.Canvas)

	// Function to dynamically resize the canvas
	resizeCanvas := js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		fmt.Println("Resizing canvas")
		// window := js.Global().Get("window")
		body := js.Global().Get("document").Get("body")

		bodyWidth := body.Get("clientWidth").Int()
		bodyHeight := body.Get("clientHeight").Int()
		fmt.Println("Body width:", bodyWidth)
		fmt.Println("Body height:", bodyHeight)

		var width, height int
		if gameWindow.FillMaxSize {
			if gameWindow.AspectRatio != nil {
				if float64(bodyWidth)/float64(bodyHeight) > *gameWindow.AspectRatio {
					height = bodyHeight
					width = int(float64(height) * *gameWindow.AspectRatio)
				} else {
					width = bodyWidth
					height = int(float64(width) / *gameWindow.AspectRatio)
				}
			} else {
				width = bodyWidth
				height = bodyHeight
			}
		} else {
			if gameWindow.AspectRatio != nil {
				width = 800
				height = int(800 / *gameWindow.AspectRatio)
			} else {
				width = *gameWindow.Width
				height = *gameWindow.Height
			}
		}

		gameWindow.Canvas.Set("width", width)
		gameWindow.Canvas.Set("height", height)

		// Update the WebGL viewport
		gl := gameWindow.Canvas.Call("getContext", "webgl")
		if !gl.IsNull() {
			fmt.Println("Setting viewport")
			gl.Call("viewport", 0, 0, width, height)
			gl.Call("clearColor", 0.0, 0.0, 0.0, 1.0)
			gl.Call("clear", gl.Get("COLOR_BUFFER_BIT"))
		}

		return nil
	})

	// Attach the resize event listener
	js.Global().Get("window").Call("addEventListener", "resize", resizeCanvas)

	// Trigger the resize function initially
	resizeCanvas.Invoke(js.Null(), nil)

	// Get the WebGL rendering context
	gl := gameWindow.Canvas.Call("getContext", "webgl")
	if gl.IsNull() {
		fmt.Println("WebGL not supported")
		return
	}

	// // Set the clear color and clear the canvas
	// gl.Call("clearColor", 0.0, 0.0, 0.0, 1.0)
	// gl.Call("clear", gl.Get("COLOR_BUFFER_BIT"))
}
