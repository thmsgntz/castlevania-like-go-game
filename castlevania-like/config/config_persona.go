package config

import "path/filepath"

var (
	// MetroidVania-like: https://o-lobster.itch.io/platformmetroidvania-pixel-art-asset-pack

	// Hero
	HeroPngPath string

	// Ennemies
	EnnemiDir string
	EnnFlyDir string
	EnnFlyPng string
)

func init_persona(assets_dir string) {

	HeroPngPath = filepath.Join(assets_dir, "herochar sprites(new)", "herochar_spritesheet(new).png")

	EnnemiDir = filepath.Join(assets_dir, "ennemies sprites")
	EnnFlyDir = filepath.Join(EnnemiDir, "fly")
	EnnFlyPng = filepath.Join(EnnFlyDir, "fly_spritesheet.png")
}
