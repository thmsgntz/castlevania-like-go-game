package main

import (
	"fmt"
	"os"

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
	var gameUI GameUI
	var err error
	var event sdl.Event
	var personaPosition *sdl.Point

	err = sdl.Init(sdl.INIT_EVERYTHING)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Sdl.init failed :%s\n", err)
		return 1
	}
	defer sdl.Quit()

	err = gameUI.InitUI(
		WinTitle,
		sdl.WINDOWPOS_UNDEFINED,
		sdl.WINDOWPOS_UNDEFINED,
		WinWidth,
		WinHeight,
		sdl.WINDOW_SHOWN,
	)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to init.\n")
		return 1
	}
	defer gameUI.DestroyUI()

	gameUI.ClearAndSetBackground(BackgroundStart)
	gameUI.LoadPersonaTexture(PersonaPngPath)

	personaPosition = &sdl.Point{X: 20, Y: 20}
	gameUI.DrawPersona(personaPosition)
	gameUI.renderer.Present()

	// while true show
	running := true
	for running {
		if event = sdl.PollEvent(); event != nil {
			fmt.Printf("Recu: %#T %v\n", event, event)
			switch t := event.(type) {
			case *sdl.QuitEvent:
				running = false
			case *sdl.KeyboardEvent:
				// gestion des KEYS: https://github.com/veandco/go-sdl2-examples/blob/29a79b36df6da7ecbcb99360a99f9e71a3cf6413/examples/keyboard-input/keyboard-input.go

				keyCode := t.Keysym.Sym

				switch {
				case keyCode == sdl.K_ESCAPE:
					running = false

				case keyCode == sdl.K_DOWN:
					personaPosition.Y += 10

				case keyCode == sdl.K_UP:
					personaPosition.Y -= 10

				case keyCode == sdl.K_LEFT:
					personaPosition.X -= 10

				case keyCode == sdl.K_RIGHT:
					personaPosition.X += 10
				}

			default:
				fmt.Printf("\n>Event non géré %v %v\n\n", event, event.GetType())
			}
			gameUI.ClearAndSetBackground(BackgroundStart)
			gameUI.DrawPersona(personaPosition)
			gameUI.renderer.Present()
			sdl.Delay(16)
		}
	}

	// Add character

	return 0
}

func main() {
	fmt.Println("\nCalling Main from main.go..")
	runGame()

}
