package main

import (
	"fmt"
	"os"

	"github.com/veandco/go-sdl2/img"
	"github.com/veandco/go-sdl2/sdl"
)

const (
	personaHeight, personaWidth                = 64, 64
	personaTileSize                            = 16
	personaTilePositionX, personaTilePositionY = 0, 0
	personaUISize                              = 64
)

var (
	personaTilePositionOnImage = sdl.Rect{X: personaTilePositionX, Y: personaTilePositionY, W: personaTileSize, H: personaTileSize}
)

type GameUI struct {
	window          *sdl.Window
	renderer        *sdl.Renderer
	personaTexture  *sdl.Texture
	personaPosition *sdl.Point
}

func (g *GameUI) InitUI(title string, x int32, y int32, w int32, h int32, flags uint32) error {
	var err error
	g.window, err = CreateSdlWindow(title, x, y, w, h, flags)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to create window: %s\n", err)
		return err
	}

	g.renderer, err = CreateSdlRenderer(g.window, -1, sdl.RENDERER_ACCELERATED)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to create renderer: %s\n", err)
		return err
	}

	g.personaPosition = &sdl.Point{X: 0, Y: 0}

	return nil
}

func (g *GameUI) ClearAndSetBackground(imgPath string) error {
	var texture *sdl.Texture
	var surface *sdl.Surface
	var err error

	// Read image and Copy it in background
	surface, err = img.Load(imgPath)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to load PNG: %s\n", err)
		return err
	}
	defer surface.Free()

	texture, err = g.renderer.CreateTextureFromSurface(surface)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to create texture: %s\n", err)
		return err
	}
	defer texture.Destroy()

	g.renderer.Clear()
	g.renderer.Copy(texture, nil, nil)

	return nil
}

func (g *GameUI) LoadPersonaTexture(imgPath string) error {
	var surface *sdl.Surface
	var err error
	surface, err = img.Load(imgPath)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to load Persona surface PNG: %s\n", err)
		return err
	}
	defer surface.Free()

	g.personaTexture, err = g.renderer.CreateTextureFromSurface(surface)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to create texturePersona: %s\n", err)
		return err
	}

	return nil
}

func (g *GameUI) DrawPersona(position *sdl.Point) {
	g.renderer.Copy(g.personaTexture, &personaTilePositionOnImage, &sdl.Rect{X: position.X, Y: position.Y, W: personaUISize, H: personaUISize})
}

func (g *GameUI) DrawTexture(texture *sdl.Texture, src *sdl.Rect, dst *sdl.Rect) {
	g.renderer.Copy(texture, src, dst)
}

func (g *GameUI) DestroyUI() {
	fmt.Printf("Destroying UI..")
	g.personaTexture.Destroy()
	g.renderer.Destroy()
	g.window.Destroy()
}
