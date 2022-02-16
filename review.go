package main

import (
	"math/rand"
	"sort"
)

func runReview(actors Actors) {
	sort.Sort(actors)

	//	outputActors("red", actors)
	round := 1
	numAlive := actors.Len()
	playerAction := 0
	for {
		for x := 0; x < actors.Len(); x++ {
			actors[x].Evasion = 0
		}
		Output("green", "\nCode Review patchset ", round, " begins...")
		outputActors("green", actors)
		for x := 0; x < actors.Len(); x++ {
			if actors[x].Morale <= 0 {
				continue
			}
			playerAction = 0
			if !actors[x].Npc {
				Output("blue", "What Do you want to do?")
				Output("blue", "\t1 - Run")
				Output("blue", "\t2 - Evade")
				Output("blue", "\t3 - Attack")
				UserInput(&playerAction)
			}
			if playerAction == 2 {
				actors[x].Evasion = rand.Intn(15)
				Output("green", "Evasion set to:", actors[x].Evasion)
			}
			tgt := selectTarget(actors, x)
			if tgt != -1 {
				// Output("red", "player: ", x, ", target: ", tgt)
				attp1 := actors[x].Attack() - actors[tgt].Evasion
				if attp1 < 0 {
					attp1 = 0
				}
				actors[tgt].Morale = actors[tgt].Morale - attp1
				if actors[tgt].Morale <= 0 {
					numAlive--
				}
				Output("green", actors[x].Name+" uses ", Tactics[actors[x].Tactic].Name, " to affect Morale by ", attp1, ".")
			}
		}
		if reviewEnded(actors) || playerAction == 1 {
			break
		} else {
			//			outputActors("green", actors)
			round++
		}
	}
	//	outputActors("black", actors)
	Output("green", "Code Review is over...")
	for x := 0; x < actors.Len(); x++ {
		if actors[x].Morale > 0 {
			Output("blue", actors[x].Name+" is still working on the change!!!")
		}
	}
}

func outputActors(color string, actors Actors) {
	for x := 0; x < actors.Len(); x++ {
		actors[x].Output(color)
	}
}

// This is a little silly, because npcs can only ever target the player,
// and there are currently only ever two participants in a code review.
// But it might get more useful if we expand.
func selectTarget(actors []Actor, selectorIndex int) int {
	y := selectorIndex
	for {
		y = y + 1
		if y >= len(actors) {
			y = 0
		}
		if (actors[y].Npc != actors[selectorIndex].Npc) && actors[y].Morale > 0 {
			return y
		}
		if y == selectorIndex {
			return -1
		}
	}
	return -1
}

func reviewEnded(actors []Actor) bool {
	count := make([]int, 2)
	count[0] = 0
	count[1] = 0
	for _, pla := range actors {
		if pla.Morale > 0 {
			if pla.Npc == false {
				count[0]++
			} else {
				count[1]++
			}
		}
	}
	if count[0] == 0 || count[1] == 0 {
		return true
	} else {
		return false
	}
}
