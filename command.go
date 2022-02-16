package main

import (
	"os"
	"strings"
)

func ProcessCommands(player *Actor, input string) {
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
	loc := LocationMap[player.CurrentLocation]
	switch command {
	case "goto":
		if loc.CanGoTo(strings.ToLower(param1)) {
			locName, err := FindLocationName(strings.ToLower(param1))
			if err != nil {
				Output("red", "Can't go to "+param1+" from here.")
			} else {
				player.CurrentLocation = locName
			}
		} else {
			Output("red", "Can't go to "+param1+" from here.")
		}
	case "help":
		Output("blue", "Commands:")
		Output("blue", "\tgoto <Location Name> - Move to the new location")
		Output("blue", "\thelp View this help screen")
		Output("blue", "\tquit Abandon your change and exit the game")
		Output("blue", "\n\n")
	case "quit":
		Output("green", "You have abandoned your patch. Goodbye...")
		os.Exit(0)
	default:
	}
}
