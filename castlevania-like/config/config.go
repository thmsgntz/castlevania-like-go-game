package config

import (
	"fmt"
	"os"
	"path/filepath"
)

var (
	ExecutableName string
	BinDir         string
	RootDir        string
	SourceDir      string
	AssetsDir      string
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
	AssetsDir = filepath.Join(RootDir, "ressources", "assets", "metroidvania")

	// INIT UI variables
	init_ui(AssetsDir)

	// INIT PERSONA variables
	init_persona(AssetsDir)
}
