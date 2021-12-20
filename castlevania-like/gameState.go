package main

import (
	"github.com/veandco/go-sdl2/sdl"
)

type GameState struct {
	running bool
	hero    *Persona
}

func (state *GameState) Init() error {
	state.running = true

	// Init Hero
	state.hero = &Persona{
		texture: nil,
		tileRect: &sdl.Rect{
			X: 0,
			Y: 5 * HeroTileSize,
			W: HeroTileSize,
			H: HeroTileSize,
		},

		tileIdlingPos: &sdl.Point{X: 0, Y: 5},
		tileIdleNbImg: 4,

		tileRunningPos:   &sdl.Point{X: 0, Y: 1},
		tileRunningNbImg: 6,

		name:        "HOTG",
		uiSize:      HeroUISize,
		typePersona: PERSONA_TYPE_HERO,

		positionUI: &sdl.Point{X: 10, Y: 10},
		pdv:        10,
		posture:    &POSTURE_CHAR_IDLE,

		direction: &DIRECTION_DROITE,
	}

	return nil
}

func (state *GameState) Destroy() {
	state.hero.texture.Destroy()
}

func (state *GameState) Move(shift *sdl.Point) error {
	return state.hero.MovePersona(shift)
}

func (state *GameState) StopMoving() {
	if state.hero.posture == &POSTURE_CHAR_RUNNING {
		state.hero.posture = &POSTURE_CHAR_IDLE
		state.hero.tileRect.Y = state.hero.tileIdlingPos.Y * HeroTileSize
		state.hero.tileRect.X = state.hero.tileIdlingPos.X * HeroTileSize
	}
}
