package main

// Game top-level structure for the game
// Not doing much with this yet, but I feel like I eventually may
type Game struct {
	Player Actor
}

func (g *Game) Play() {
	// We should prompt people for their name. They can pick Jimmy if they want to.
	g.Player = *new(Actor)
	g.Player.Name = "Jimmy Wales"
	g.Player.Morale = 100
	g.Player.Tactics = []int{1, 2, 3, 4, 5, 6}
	g.Player.CurrentLocation = "CommandLine"

	Output("blue", Messages["welcome"])
	for {
		Output("blue", LocationMap[g.Player.CurrentLocation].Description)

		// We really shouldn't process an event unless location has changed.
		// Otherwise you can stay in AFK forever and get crazy morale by hitting Return over and over
		g.ProcessEvents(LocationMap[g.Player.CurrentLocation].Events)
		if g.Player.Morale <= 0 {
			Output("white", "You have given up hope on your change. Game over!!!")
			return
		} else {
			Output("white", "You are still working on your change.")
		}

		Output("blue", "Morale:", g.Player.Morale)
		Output("green", "You can go to these places:")
		for _, loc := range LocationMap[g.Player.CurrentLocation].Transitions {
			Outputf("green", "\t%s", loc)
		}
		cmd := UserInputln()
		ProcessCommands(&g.Player, cmd)
	}
}

func (g *Game) ProcessEvents(events []string) {
	for _, evtName := range events {
		g.Player.Morale += Events[evtName].ProcessEvent(&g.Player)
	}
}
