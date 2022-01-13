package config

import "path/filepath"

const (
	WinTitle            string = "CastlevaniaLike!"
	WinWidth, WinHeight int32  = 800, 600
)

var (
	// MetroidVania-like: https://o-lobster.itch.io/platformmetroidvania-pixel-art-asset-pack
	BackgroundStart string
)

func init_ui(assets_dir string) {
	BackgroundStart = filepath.Join(assets_dir, "tiles_and_background_foreground_(new)", "background.png")
}
