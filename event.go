package main

import (
	"math/rand"
	"time"
)

type Event struct {
	Type        string
	Chance      int
	Description string
	Morale      int
	Evt         string
}

func (e *Event) ProcessEvent(player *Actor) int {
	s1 := rand.NewSource(time.Now().UnixNano())
	r1 := rand.New(s1)
	moraleAdjustment := 0
	if e.Chance >= r1.Intn(100) {
		if e.Type == "CodeReview" {
			// Generate reviewer
			reviewer := new(Actor)
			*reviewer = *Reviewers[1+rand.Intn(len(Reviewers)-1)]
			reviewer.Npc = true
			Output("green", "\tA "+reviewer.Name+" reviews your code.")

			actors := Actors{*reviewer, *player}
			moraleAdjustment = runReview(actors)
		} else {
			moraleAdjustment = e.Morale
			Output("green", "\t"+e.Description+" affecting your morale by ", e.Morale)
			if e.Evt != "" {
				moraleAdjustment += Events[e.Evt].ProcessEvent(player)
			}
		}
	}
	return moraleAdjustment
}
