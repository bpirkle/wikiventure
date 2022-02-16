package main

// Actor Used for both the player and all npcs
// By convention, if a variable name is "player", it refers
// to the main player character and not an npc.
type Actor struct {
	Name       string
	Morale     int
	Evasion    int
	Initiative int
	Tactic     int
	Npc        bool
	Items      []int

	CurrentLocation string
}

func (a *Actor) Equip(t int) {
	a.Tactic = t
}

func (a *Actor) Attack() int {
	return Tactics[a.Tactic].Use()
}

type Actors []Actor

func (slice Actors) Len() int {
	return len(slice)
}

func (slice Actors) Less(i, j int) bool {
	return slice[i].Initiative > slice[j].Initiative // Sort descending
}

func (slice Actors) Swap(i, j int) {
	slice[i], slice[j] = slice[j], slice[i]
}

func (a *Actor) Output(color string) {
	Output(color, a.Name, ": Morale ", a.Morale)
}
