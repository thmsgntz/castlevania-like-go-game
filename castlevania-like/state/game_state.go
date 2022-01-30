package state

import (
	"castlevania-like-go-game/castlevania-like/config"
	"castlevania-like-go-game/castlevania-like/persona"
	"fmt"

	"github.com/veandco/go-sdl2/sdl"
)

type GameState struct {
	running  bool
	hero     *persona.Persona
	ennemies []*persona.Persona
}

func (state *GameState) Init() error {
	state.running = true

	// Init Hero
	state.hero = persona.CreateHero(&sdl.Point{X: 10, Y: 10}, "Hero 1")
	state.SetMapLevel(1)

	return nil
}

func (state *GameState) AddEnnemy(p *persona.Persona) error {

	// Ajout d'un ennemi dans le tableau des ennemies
	// Checker taille / capacity ??
	state.ennemies = append(state.ennemies, p)

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

func (state *GameState) Enemies() []*persona.Persona {
	return state.ennemies
}

func (state *GameState) SetMapLevel(levelNb int) (err error) {
	switch levelNb {
	case 1:
		{
			LoadLevelFromFile(config.MapLevel1Path)
		}
	default:
		{
			return fmt.Errorf("Level Number does not exist")
		}
	}
	return err
}
