package main

import (
	"fmt"
	"os"

	"github.com/veandco/go-sdl2/sdl"

	// "github.com/veandco/go-sdl2/img"
	// "github.com/veandco/go-sdl2/sdl"
	"castlevania-like-go-game/castlevania-like/UI"
	"castlevania-like-go-game/castlevania-like/state"
	_ "castlevania-like-go-game/utils"
)

func runGame() int {
	var gameUI UI.GameUI
	var gState state.GameState
	var err error
	var event sdl.Event

	err = sdl.Init(sdl.INIT_EVERYTHING)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Sdl.init failed :%s\n", err)
		return 1
	}
	defer sdl.Quit()

	err = gState.Init()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to init the state! %s.\n", err)
		return 1
	}
	defer gState.Destroy()

	err = gameUI.InitUI(
		state.WinTitle,
		sdl.WINDOWPOS_UNDEFINED,
		sdl.WINDOWPOS_UNDEFINED,
		state.WinWidth,
		state.WinHeight,
		sdl.WINDOW_SHOWN,
	)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to init gameUI! %s.\n", err)
		return 1
	}
	defer gameUI.DestroyUI()

	gameUI.SetBackground(state.BackgroundStart)
	gameUI.SetPersonaTexture(gState.Hero(), state.HeroPngPath)

	// while true show
	running := true
	for running {
		if event = sdl.PollEvent(); event != nil {
			// fmt.Printf("Recu: %#T %v\n", event, event)
			switch t := event.(type) {
			case *sdl.QuitEvent:
				running = false
			case *sdl.KeyboardEvent:
				// gestion des KEYS: https://github.com/veandco/go-sdl2-examples/blob/29a79b36df6da7ecbcb99360a99f9e71a3cf6413/examples/keyboard-input/keyboard-input.go

				if t.State == sdl.RELEASED {
					gState.StopMoving()
					break
				}

				keyCode := t.Keysym.Sym

				switch {
				case keyCode == sdl.K_ESCAPE:
					running = false

				case keyCode == sdl.K_DOWN:
					gState.Move(&sdl.Point{X: 0, Y: +10})

				case keyCode == sdl.K_UP:
					gState.Move(&sdl.Point{X: 0, Y: -10})

				case keyCode == sdl.K_LEFT:
					gState.Move(&sdl.Point{X: -10, Y: 0})

				case keyCode == sdl.K_RIGHT:
					gState.Move(&sdl.Point{X: 10, Y: 0})
				}

			default:
				// fmt.Printf("\n>Event non géré %v %v\n\n", event, event.GetType())
			}

			gameUI.Update(gState.Hero())
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
