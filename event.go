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
	if e.Chance >= r1.Intn(100) {
		if e.Type == "CodeReview" {
			// Generate reviewer
			reviewer := new(Actor)
			*reviewer = *Reviewers[1+rand.Intn(len(Reviewers)-1)]
			reviewer.Npc = true
			Output("green", "A "+reviewer.Name+" reviews your code.")

			actors := Actors{*reviewer, *player}
			runReview(actors)
		} else {
			Output("green", "\t"+e.Description)
			if e.Evt != "" {
				Output("green", "\t"+Events[e.Evt].Description)
				e.Morale = e.Morale + Events[e.Evt].ProcessEvent(player)
			}
		}
		return e.Morale
	}
	return 0
}
