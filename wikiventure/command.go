package wikiventure

import (
	"os"
	"strings"
)

func ProcessCommands(game *Game, input string) {
	Output("yellow", "======================================================================")
	tokens := strings.Fields(input)
	if len(tokens) == 0 {
		Output("red", "No command received.")
		return
	}
	command := strings.ToLower(tokens[0])
	param1 := ""
	if len(tokens) > 1 {
		param1 = tokens[1]
	}
	switch command {
	case "goto":
		loc := LocationMap[game.Player.CurrentLocation]
		locName, err := FindLocationName(strings.ToLower(param1))
		if err != nil {
			Output("red", err)
		} else if loc.CanGoTo(strings.ToLower(locName)) {
			game.Player.CurrentLocation = locName
		} else {
			Output("red", "Can't go to "+param1+" from here.")
		}
	case "color":
		game.setColorScheme(param1)
	case "help":
		Output("blue", "Commands:")
		Output("blue", "\tgoto <Location Name> - Move to the new location")
		Output("blue", "\tcolor <dark|light|none> - Set text colors for a dark or light terminal background (or just disable colors)")
		Output("blue", "\thelp View this help screen")
		Output("blue", "\tquit Abandon your change and exit the game")
		Output("blue", "\n\n")
	case "quit":
		Output("green", "You have abandoned your change. Goodbye...")
		os.Exit(0)
	default:
	}
}
