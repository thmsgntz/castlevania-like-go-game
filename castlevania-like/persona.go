package main

import (
	"fmt"

	"github.com/veandco/go-sdl2/sdl"
)

type PersonaType int

var (
	PERSONA_TYPE_HERO  PersonaType = 0
	PERSONA_TYPE_ENEMY PersonaType = 1
)

type Posture int

var (
	// Posture for Persona
	POSTURE_CHAR_IDLE    Posture = 0
	POSTURE_CHAR_RUNNING Posture = 1
)

const (
	HeroHeight, HeroWidth                        = 64, 64
	HeroTileSize                                 = 16
	HeroUISize                                   = 64
	HeroTileIdlePositionX, HeroTileIdlePositionY = 0, 5
)

type Persona struct {
	texture  *sdl.Texture
	tileRect *sdl.Rect // sdl.Rect utilisé pour render.Copy (src). Selectionne la tile

	tileIdlingPos *sdl.Point // Index X,Y dans l'image (première élément)
	tileIdleNbImg int        // Nombre d'image pour l'animation Idle

	tileRunningPos   *sdl.Point // Index X,Y dans l'image (première élément)
	tileRunningNbImg int        // Nombre d'image pour l'animation Running

	uiSize int32

	name        string
	typePersona PersonaType

	pdv        uint
	positionUI *sdl.Point
	posture    *Posture
}

func (persona *Persona) MovePersona(shift *sdl.Point) error {
	// Access global variable WinWidth, WinHeight from confi.go

	fmt.Printf("%v %#v\n", *persona.posture, persona.tileRect)

	// Tile Animation
	if persona.posture == &POSTURE_CHAR_RUNNING {
		// already running, load next running animation
		if persona.tileRunningPos.X+1 == int32(persona.tileRunningNbImg) {
			persona.tileRunningPos.X = 0
			persona.tileRect.X = 0
		} else {
			persona.tileRunningPos.X += 1
			persona.tileRect.X += HeroTileSize
		}
	} else {
		// was idling. Start running.
		persona.posture = &POSTURE_CHAR_RUNNING
		persona.tileRunningPos.X = 0
		persona.tileRect.X = 0
	}

	// Rect adjusting
	persona.tileRect.X = persona.tileRunningPos.X * HeroTileSize
	persona.tileRect.Y = persona.tileRunningPos.Y * HeroTileSize

	// Position on GUI
	var shiftTmp int32 = persona.positionUI.X + shift.X
	switch {
	case shiftTmp < 0:
		persona.positionUI.X = 0
	case shiftTmp > WinWidth:
		persona.positionUI.X = WinWidth
	default:
		persona.positionUI.X = shiftTmp
	}

	shiftTmp = persona.positionUI.Y + shift.Y
	switch {
	case shiftTmp < 0:
		persona.positionUI.Y = 0
	case shiftTmp > WinHeight:
		persona.positionUI.Y = WinHeight
	default:
		persona.positionUI.Y = shiftTmp
	}

	fmt.Printf("%v Pos:{%v, %v}\n", persona.name, persona.positionUI.X, persona.positionUI.Y)

	return nil
}