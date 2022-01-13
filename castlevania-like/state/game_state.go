package state

import (
	"castlevania-like-go-game/castlevania-like/persona"

	"github.com/veandco/go-sdl2/sdl"
)

type GameState struct {
	running bool
	hero    *persona.Persona
	// ennemies []*Persona
}

func (state *GameState) Init() error {
	state.running = true

	// Init Hero
	state.hero = persona.CreateHero(&sdl.Point{X: 10, Y: 10})

	return nil
}

func (state *GameState) AddEnnemy(enn_type *persona.PersonaType) error {
	/* if enn_type == &PERSONA_TYPE_E_FLY {

	} */
	return nil
}

func (state *GameState) Destroy() {
	state.hero.Texture.Destroy()
}

func (state *GameState) Move(shift *sdl.Point) error {
	return state.hero.MovePersona(shift)
}

func (state *GameState) StopMoving() error {
	return state.hero.StopMoving()
}

func (state *GameState) Hero() *persona.Persona {
	return state.hero
}
