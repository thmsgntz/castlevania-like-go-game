package UI

import (
	"fmt"
	"os"

	"github.com/veandco/go-sdl2/sdl"
)

func CreateSdlWindow(title string, x int32, y int32, w int32, h int32, flags uint32) (window *sdl.Window, err error) {
	window, err = sdl.CreateWindow(title, x, y, w, h, flags)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to create window: %s\n", err)
		return
	}

	return
}

func CreateSdlRenderer(window *sdl.Window, index int, flags uint32) (renderer *sdl.Renderer, err error) {
	renderer, err = sdl.CreateRenderer(window, index, flags)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to create renderer: %s\n", err)
		return
	}

	return
}
