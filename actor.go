package main

// Actor Used for both the player and all npcs
// By convention, if a variable name is "player", it refers
// to the main player character and not an npc.
type Actor struct {
	Name    string
	Morale  int
	Tactics []int
	Npc     bool

	CurrentLocation string
}

func (a *Actor) Act(action int) (int, *Tactic) {
	tactic := a.Tactics[action]
	return Tactics[tactic].Use(), Tactics[tactic]
}

type Actors []Actor

func (slice Actors) Len() int {
	return len(slice)
}

func (a *Actor) Output(color string) {
	Output(color, a.Name, ": Morale ", a.Morale)
}
