package main

import (
	"fmt"

	"github.com/veandco/go-sdl2/sdl"
)

type Posture string

var (
	// Posture for Persona
	CHAR_RUNNING Posture = "RUNNING"
	CHAR_IDLE    Posture = "IDLING"
)

const (
	HeroHeight, HeroWidth                = 64, 64
	HeroTileSize                         = 16
	HeroTilePositionX, HeroTilePositionY = 0, 0
	HeroUISize                           = 64
)

var (
	HeroTilePositionOnImage = sdl.Rect{X: HeroTilePositionX, Y: HeroTilePositionY, W: HeroTileSize, H: HeroTileSize}
)

type GameState struct {
	running bool
	hero    *Persona
}

type Persona struct {
	texture  *sdl.Texture
	position *sdl.Point

	uiSize int32

	name string

	pdv     uint
	posture *Posture
}

func (state *GameState) Init() error {
	state.running = true

	// Init Hero
	state.hero = &Persona{
		name:     "HOTG",
		uiSize:   HeroUISize,
		texture:  nil,
		position: &sdl.Point{X: 10, Y: 10},
		pdv:      10,
		posture:  &CHAR_IDLE,
	}

	return nil
}

func (state *GameState) Destroy() {
	state.hero.texture.Destroy()
}

func (state *GameState) Move(shift *sdl.Point) error {
	return state.hero.MovePersona(shift)
}

func (persona *Persona) MovePersona(shift *sdl.Point) error {
	// Access global variable WinWidth, WinHeight from confi.go
	var shiftTmp int32 = persona.position.X + shift.X

	switch {
	case shiftTmp < 0:
		persona.position.X = 0
	case shiftTmp > WinWidth:
		persona.position.X = WinWidth
	default:
		persona.position.X = shiftTmp
	}

	shiftTmp = persona.position.Y + shift.Y
	switch {
	case shiftTmp < 0:
		persona.position.Y = 0
	case shiftTmp > WinHeight:
		persona.position.Y = WinHeight
	default:
		persona.position.Y = shiftTmp
	}

	fmt.Printf("%v Pos:{%v, %v}\n", persona.name, persona.position.X, persona.position.Y)

	return nil
}
