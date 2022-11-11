package main

import (
	"github.com/xm-tech/guess-game/internal"
)

func main() {
	internal.G = internal.NewGame()
	internal.G.Run()
}
