package main

import (
	"math/rand"
)

func runReview(actors Actors) {
	consensus := 0
	round := 1
	action := 0
	for {
		Output("green", "\nCode Review patchset ", round, " begins...")
		outputActors("green", actors)
		for x := 0; x < actors.Len(); x++ {
			if actors[x].Morale <= 0 {
				continue
			}
			if !actors[x].Npc {
				Output("blue", "What Do you want to do?")
				for option := 0; option < len(actors[x].Actions); option++ {
					Output("blue", "\t", option+1, " - ", Actions[actors[x].Actions[option]].Name)
				}
				UserInput(&action)
				action--
			} else {
				action = rand.Intn(len(actors[x].Actions))
			}
			tgt := selectTarget(actors, x)
			if tgt != -1 {
				var effect, actionName = actors[x].Act(action)
				actors[tgt].Morale = actors[tgt].Morale + effect
				if effect < 0 {
					consensus -= effect
				} else {
					consensus += effect
				}
				if consensus > 100 {
					consensus = 100
				}

				Output("green", actors[x].Name+" uses ", actionName, " to affect Morale by ", effect, ".")
				Output("green", "Consensus is at ", consensus, "%")
			}
		}
		if isReviewEnded(actors, consensus) {
			break
		} else {
			round++
		}
	}

	Output("green", "Code Review is over.")
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

func isReviewEnded(actors []Actor, consensus int) bool {
	if consensus >= 100 {
		return true
	}

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
