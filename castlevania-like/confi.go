package main

import (
	"fmt"
	"os"
	"path/filepath"
)

const (
	WinTitle            string = "CastlevaniaLike!"
	WinWidth, WinHeight int32  = 800, 600
)

var (
	ExecutableName string
	BinDir         string
	RootDir        string
	SourceDir      string
	AssetsDir      string

	// MetroidVania-like: https://o-lobster.itch.io/platformmetroidvania-pixel-art-asset-pack
	MetroidVaniaDir string
	BackgroundStart string
	PersonaPngPath  string
)

func init() {
	fmt.Println("\nCalling Init from confi.go..")
	var err error
	ExecutableName, err = os.Executable()
	if err != nil {
		panic(err)
	}

	BinDir, err = filepath.Abs(filepath.Dir(ExecutableName))
	if err != nil {
		panic(err)
	}

	RootDir = filepath.Join(BinDir, "..")
	SourceDir = filepath.Join(RootDir, "src")
	AssetsDir = filepath.Join(RootDir, "ressources", "assets")

	MetroidVaniaDir = filepath.Join(AssetsDir, "metroidvania")
	BackgroundStart = filepath.Join(MetroidVaniaDir, "tiles and background_foreground (new)", "background.png")
	PersonaPngPath = filepath.Join(MetroidVaniaDir, "herochar sprites(new)", "herochar_spritesheet(new).png")
}
