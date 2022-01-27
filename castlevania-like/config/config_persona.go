package config

import "path/filepath"

var (
	// MetroidVania-like: https://o-lobster.itch.io/platformmetroidvania-pixel-art-asset-pack

	// Hero
	HeroPngPath string

	// Enemies
	EnemiDir string
	EnFlyDir string
	EnFlyPng string
)

func init_persona(assets_dir string) {

	HeroPngPath = filepath.Join(assets_dir, "herochar sprites(new)", "herochar_spritesheet(new).png")

	EnemiDir = filepath.Join(assets_dir, "enemies sprites")
	EnFlyDir = filepath.Join(EnemiDir, "fly")
	EnFlyPng = filepath.Join(EnFlyDir, "fly_spritesheet.png")
}
