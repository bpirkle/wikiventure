package main

import (
	"wikiventure/wikiventure"
)

func main() {
	var game = *new(wikiventure.Game)
	game.Play()
}
