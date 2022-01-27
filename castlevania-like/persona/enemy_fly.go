package persona

import "github.com/veandco/go-sdl2/sdl"

func CreateEnnemiFly(location *sdl.Point, name string) *Persona {
	persona := &Persona{
		Texture: nil,
		TileRect: &sdl.Rect{
			X: 0,
			Y: 1 * 8,
			W: 8,
			H: 8,
		},

		TileIdlingPos: &sdl.Point{X: 0, Y: 1},
		TileIdleNbImg: 3,

		TileRunningPos:   &sdl.Point{X: 0, Y: 3},
		TileRunningNbImg: 3,

		Name:        name,
		UiSize:      64,
		TypePersona: PERSONA_TYPE_ENEMY_FLY,

		PositionUI: location,
		Pdv:        3,
		Posture:    &POSTURE_CHAR_IDLE,

		Direction: &DIRECTION_DROITE,
	}

	return persona
}
