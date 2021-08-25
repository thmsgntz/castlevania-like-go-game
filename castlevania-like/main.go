package main

import (
	"fmt"
	"os"

	"github.com/veandco/go-sdl2/img"
	"github.com/veandco/go-sdl2/sdl"

	// "github.com/veandco/go-sdl2/img"
	// "github.com/veandco/go-sdl2/sdl"
	_ "castlevania-like-go-game/utils"
)

/* func createBackground() int {
	var window *sdl.Window
	var renderer *sdl.Renderer
	var texture *sdl.Texture
	var src, dst sdl.Rect
	var err error

	var imageName string = filepath.Join(MetroidVaniaDir, "tiles and background_foreground (new)", "background.png")

	window, err = CreateSdlWindow(WinTitle, sdl.WINDOWPOS_UNDEFINED, sdl.WINDOWPOS_UNDEFINED,
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

	src = sdl.Rect{X: 0, Y: 0, W: 512, H: 512}
	dst = sdl.Rect{X: 100, Y: 50, W: 512, H: 512}

	renderer.Clear()
	renderer.SetDrawColor(255, 0, 0, 255)
	renderer.FillRect(&sdl.Rect{X: 0, Y: 0, W: int32(WinWidth), H: int32(WinHeight)})
	renderer.Copy(texture, &src, &dst)
	renderer.Present()

	sdl.PollEvent()
	sdl.Delay(10000)

	return 0
} */

func runGame() int {
	var window *sdl.Window
	var renderer *sdl.Renderer
	var texture *sdl.Texture
	var imgSurface *sdl.Surface
	var err error
	var src, dst sdl.Rect
	var event sdl.Event

	// Create Windows & Rendered
	window, err = CreateSdlWindow(
		WinTitle,
		sdl.WINDOWPOS_UNDEFINED,
		sdl.WINDOWPOS_UNDEFINED,
		WinWidth,
		WinHeight,
		sdl.WINDOW_SHOWN,
	)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to create window: %s\n", err)
		return 1
	}
	defer window.Destroy()

	renderer, err = CreateSdlRenderer(window, -1, sdl.RENDERER_ACCELERATED)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to create renderer: %s\n", err)
		return 2
	}
	defer renderer.Destroy()

	// Read image and Copy it in background
	imgSurface, err = img.Load(BackgroundStart)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to load PNG: %s\n", err)
		return 3
	}
	defer imgSurface.Free()

	texture, err = renderer.CreateTextureFromSurface(imgSurface)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to create texture: %s\n", err)
		return 4
	}
	defer texture.Destroy()

	src = sdl.Rect{X: 0, Y: 0, W: imgSurface.W, H: imgSurface.H} // square in the source image
	dst = sdl.Rect{X: 0, Y: 0, W: WinWidth, H: WinHeight}        // square in the window destination, streched if needed

	renderer.Clear()
	renderer.SetDrawColor(255, 0, 0, 255)
	renderer.FillRect(&sdl.Rect{X: 0, Y: 0, W: int32(WinWidth), H: int32(WinHeight)})
	renderer.Copy(texture, &src, &dst)
	renderer.Present()

	// while true show
	running := true
	for running {
		if event = sdl.PollEvent(); event != nil {
			fmt.Printf("%#T %v\n", event, event.GetType())
			switch t := event.(type) {
			case *sdl.QuitEvent:
				running = false
			case *sdl.KeyboardEvent:
				// gestion des KEYS: https://github.com/veandco/go-sdl2-examples/blob/29a79b36df6da7ecbcb99360a99f9e71a3cf6413/examples/keyboard-input/keyboard-input.go
				if keyCode := t.Keysym.Sym; keyCode == sdl.K_ESCAPE {
					running = false
				}
			default:
				fmt.Printf("Event non géré %v %v\n", event, event.GetType())
			}
		}
	}

	// Add character

	return 0
}

func main() {
	fmt.Println("\nCalling Main from main.go..")
	runGame()

}
