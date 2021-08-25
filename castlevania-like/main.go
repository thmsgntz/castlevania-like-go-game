package main

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/veandco/go-sdl2/img"
	"github.com/veandco/go-sdl2/sdl"

	// "github.com/veandco/go-sdl2/img"
	// "github.com/veandco/go-sdl2/sdl"
	_ "castlevania-like-go-game/utils"
)

func createBackground() int {
	var window *sdl.Window
	var renderer *sdl.Renderer
	var texture *sdl.Texture
	var src, dst sdl.Rect
	var err error

	var imageName string = filepath.Join(MetroidVaniaDir, "tiles and background_foreground (new)", "background.png")
	window, err = sdl.CreateWindow(WinTitle, sdl.WINDOWPOS_UNDEFINED, sdl.WINDOWPOS_UNDEFINED,
		WinWidth, WinHeight, sdl.WINDOW_SHOWN)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to create window: %s\n", err)
		return 1
	}
	defer window.Destroy()

	renderer, err = sdl.CreateRenderer(window, -1, sdl.RENDERER_ACCELERATED)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to create renderer: %s\n", err)
		return 2
	}
	defer renderer.Destroy()

	image, err := img.Load(imageName)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to load PNG: %s\n", err)
		return 3
	}
	defer image.Free()

	texture, err = renderer.CreateTextureFromSurface(image)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to create texture: %s\n", err)
		return 4
	}
	defer texture.Destroy()

	src = sdl.Rect{0, 0, 512, 512}
	dst = sdl.Rect{100, 50, 512, 512}

	renderer.Clear()
	renderer.SetDrawColor(255, 0, 0, 255)
	renderer.FillRect(&sdl.Rect{0, 0, int32(WinWidth), int32(WinHeight)})
	renderer.Copy(texture, &src, &dst)
	renderer.Present()

	sdl.PollEvent()
	sdl.Delay(10000)

	return 0
}

func main() {
	fmt.Println("\nCalling Main from main.go..")
	createBackground()
}
