package state

import (
	"bufio"
	"fmt"
	"os"
)

type tile rune

const (
	tile_grass = '#'
)

type Level struct {
	Map [][]tile
}

func LoadLevelFromFile(filename string) (m *Level, err error) {
	_ = tile_grass
	// https://youtu.be/Jy919y3ezOI?list=PLDZujg-VgQlZUy1iCqBbe5faZLMkA3g2x&t=1134
	// https://pkg.go.dev/os?utm_source=gopls#Open
	fd, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	defer fd.Close()

	scanner := bufio.NewScanner(fd)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}

	return m, err
}
