package UI

import (
	"fmt"
	"os"

	"github.com/veandco/go-sdl2/img"
	"github.com/veandco/go-sdl2/sdl"

	"castlevania-like-go-game/castlevania-like/persona"
)

type GameUI struct {
	window            *sdl.Window
	renderer          *sdl.Renderer
	backgroundTexture *sdl.Texture
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

	// g.personaPosition = &sdl.Point{X: 0, Y: 0}

	return nil
}

func (g *GameUI) Update(persona *persona.Persona) error {
	var err error

	err = g.renderer.Clear()
	if err != nil {
		panic(err)
	}

	err = g.DrawTexture(g.backgroundTexture, nil, nil)
	if err != nil {
		panic(err)
	}

	err = g.DrawPersona(persona)
	if err != nil {
		panic(err)
	}

	g.renderer.Present()

	return nil
}

func (g *GameUI) SetBackground(imgPath string) error {
	var surface *sdl.Surface
	var err error

	// Read image and Copy it in background
	surface, err = img.Load(imgPath)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to load PNG: %s\n", err)
		return err
	}
	defer surface.Free()

	g.backgroundTexture, err = g.renderer.CreateTextureFromSurface(surface)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to create texture: %s\n", err)
		return err
	}

	return nil
}

func (g *GameUI) SetPersonaTexture(p *persona.Persona, imgPath string) error {
	var surface *sdl.Surface
	var err error
	surface, err = img.Load(imgPath)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to load Persona surface PNG: %s\n", err)
		return err
	}
	defer surface.Free()

	p.Texture, err = g.renderer.CreateTextureFromSurface(surface)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to create texturePersona: %s\n", err)
		return err
	}

	return nil
}

func (g *GameUI) DrawPersona(p *persona.Persona) error {
	// https://wiki.libsdl.org/SDL_RenderCopyEx

	var flip = sdl.FLIP_NONE

	if p.Direction == &persona.DIRECTION_GAUCHE {
		flip = sdl.FLIP_HORIZONTAL
	}

	g.renderer.CopyEx(
		p.Texture,  // texture
		p.TileRect, // srcrect
		&sdl.Rect{ // dstrect
			X: p.PositionUI.X,
			Y: p.PositionUI.Y,
			W: p.UiSize,
			H: p.UiSize},
		0,    // angle
		nil,  // center of rotation
		flip, // flip
	)

	return nil
}

func (g *GameUI) DrawTexture(texture *sdl.Texture, src *sdl.Rect, dst *sdl.Rect) error {
	return g.renderer.Copy(texture, src, dst)
}

func (g *GameUI) DestroyUI() {
	fmt.Printf("Destroying UI..")
	g.backgroundTexture.Destroy()
	g.renderer.Destroy()
	g.window.Destroy()
}
