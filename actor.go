package main

// Actor Used for both the player and all npcs
// By convention, if a variable name is "player", it refers
// to the main player character and not an npc.
type Actor struct {
	Name    string
	Morale  int
	Actions []int
	Npc     bool

	CurrentLocation string
}

func (a *Actor) Act(actionOption int) (int, string) {
	action := a.Actions[actionOption]
	return Actions[action].Use(), Actions[action].Name
}

type Actors []Actor

func (slice Actors) Len() int {
	return len(slice)
}

func (a *Actor) Output(color string) {
	Output(color, "\t", a.Name, ": Morale ", a.Morale)
}
