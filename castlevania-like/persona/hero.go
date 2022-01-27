package persona

import "github.com/veandco/go-sdl2/sdl"

func CreateHero(location *sdl.Point, name string) *Persona {
	persona := &Persona{
		Texture: nil,
		TileRect: &sdl.Rect{
			X: 0,
			Y: 5 * HeroTileSize,
			W: HeroTileSize,
			H: HeroTileSize,
		},

		TileIdlingPos: &sdl.Point{X: 0, Y: 5},
		TileIdleNbImg: 4,

		TileRunningPos:   &sdl.Point{X: 0, Y: 1},
		TileRunningNbImg: 6,

		Name:        name,
		UiSize:      HeroUISize,
		TypePersona: PERSONA_TYPE_HERO,

		PositionUI: location,
		Pdv:        10,
		Posture:    &POSTURE_CHAR_IDLE,

		Direction: &DIRECTION_DROITE,
	}

	return persona
}
