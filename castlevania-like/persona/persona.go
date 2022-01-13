package persona

/*

Personas are living entities that can move, interact, live and die

*/

import (
	"castlevania-like-go-game/castlevania-like/config"
	"fmt"

	"github.com/veandco/go-sdl2/sdl"
)

type PersonaType int

var (
	PERSONA_TYPE_HERO   PersonaType = 0
	PERSONA_TYPE_ENNEMY PersonaType = 1
	PERSONA_TYPE_E_FLY  PersonaType = 2
)

type Posture int

var (
	// Posture for Persona
	POSTURE_CHAR_IDLE    Posture = 0
	POSTURE_CHAR_RUNNING Posture = 1
)

type Direction int

var (
	DIRECTION_DROITE Direction = 0
	DIRECTION_GAUCHE Direction = 1
	DIRECTION_HAUT   Direction = 2
	DIRECTION_BAS    Direction = 3
)

const (
	HeroHeight, HeroWidth                        = 64, 64
	HeroTileSize                                 = 16
	HeroUISize                                   = 64
	HeroTileIdlePositionX, HeroTileIdlePositionY = 0, 5
)

type Persona struct {
	Texture  *sdl.Texture
	TileRect *sdl.Rect // sdl.Rect utilisé pour render.Copy (src). Selectionne la tile

	TileIdlingPos *sdl.Point // Index X,Y dans l'image (première élément)
	TileIdleNbImg int        // Nombre d'image pour l'animation Idle

	TileRunningPos   *sdl.Point // Index X,Y dans l'image (première élément)
	TileRunningNbImg int        // Nombre d'image pour l'animation Running

	UiSize int32

	Name        string
	TypePersona PersonaType

	Pdv        uint
	PositionUI *sdl.Point
	Posture    *Posture

	Direction *Direction // Direction (pour flipper la texture si besoin)
}

func CreateEnnemiFly(location *sdl.Point) *Persona {
	persona := &Persona{
		Texture: nil,
		TileRect: &sdl.Rect{
			X: 0,
			Y: 5 * HeroTileSize,
			W: HeroTileSize,
			H: HeroTileSize,
		},

		TileIdlingPos: &sdl.Point{X: 0, Y: 1},
		TileIdleNbImg: 3,

		TileRunningPos:   &sdl.Point{X: 0, Y: 3},
		TileRunningNbImg: 3,

		Name:        "Fly",
		UiSize:      HeroUISize,
		TypePersona: PERSONA_TYPE_E_FLY,

		PositionUI: location,
		Pdv:        10,
		Posture:    &POSTURE_CHAR_IDLE,

		Direction: &DIRECTION_DROITE,
	}

	return persona
}

func (persona *Persona) MovePersona(shift *sdl.Point) error {
	// Access global variable WinWidth, WinHeight from confi.go

	fmt.Printf("%v %#v\n", *persona.Posture, persona.TileRect)

	// Tile Animation
	if persona.Posture == &POSTURE_CHAR_RUNNING {
		// already running, load next running animation
		if persona.TileRunningPos.X+1 == int32(persona.TileRunningNbImg) {
			persona.TileRunningPos.X = 0
			persona.TileRect.X = 0
		} else {
			persona.TileRunningPos.X += 1
			persona.TileRect.X += HeroTileSize
		}
	} else {
		// was idling. Start running.
		persona.Posture = &POSTURE_CHAR_RUNNING
		persona.TileRunningPos.X = 0
		persona.TileRect.X = 0
	}

	// Rect adjusting
	persona.TileRect.X = persona.TileRunningPos.X * HeroTileSize
	persona.TileRect.Y = persona.TileRunningPos.Y * HeroTileSize

	// Position on the screen GUI on X
	if shift.X != 0 {
		var shiftTmp int32 = persona.PositionUI.X + shift.X
		switch {
		case shiftTmp < 0:
			persona.PositionUI.X = 0
		case shiftTmp > config.WinWidth:
			persona.PositionUI.X = config.WinWidth
		default:
			persona.PositionUI.X = shiftTmp
		}
	}

	if shift.Y != 0 {
		var shiftTmp = persona.PositionUI.Y + shift.Y
		switch {
		case shiftTmp < 0:
			persona.PositionUI.Y = 0
		case shiftTmp > config.WinHeight:
			persona.PositionUI.Y = config.WinHeight
		default:
			persona.PositionUI.Y = shiftTmp
		}
	}

	// Update la Direction pour flip (ou non) la texture
	persona.UpdateDirection(shift)

	fmt.Printf("%v Pos:{%v, %v}\n", persona.Name, persona.PositionUI.X, persona.PositionUI.Y)

	return nil
}

func (persona *Persona) StopMoving() error {
	if persona.Posture == &POSTURE_CHAR_RUNNING {
		persona.Posture = &POSTURE_CHAR_IDLE
		persona.TileRect.Y = persona.TileIdlingPos.Y * HeroTileSize
		persona.TileRect.X = persona.TileIdlingPos.X * HeroTileSize
	}
	return nil
}

func (persona *Persona) UpdateDirection(shift *sdl.Point) {
	// Met à jour la Direction du personnage pour que le renderer sache
	// si on a besoin de flip ou nonla texture
	// Ne gère pour l'instant que les directions horizatales

	if shift.X == 0 {
		return
	}

	if shift.X < 0 {
		persona.Direction = &DIRECTION_GAUCHE
	} else {
		persona.Direction = &DIRECTION_DROITE
	}
}
