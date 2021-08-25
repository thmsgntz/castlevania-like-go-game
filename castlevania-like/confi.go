package main

import (
	"fmt"
	"os"
	"path/filepath"
)

var ExecutableName string
var BinDir string
var RootDir string
var SourceDir string
var AssetsDir string

// MetroidVania-like: https://o-lobster.itch.io/platformmetroidvania-pixel-art-asset-pack
var MetroidVaniaDir string

var WinTitle string = "CastlevaniaLike!"
var WinWidth, WinHeight int32 = 800, 600

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
}
