package main

import (
	"math/rand"
	"time"
)

type Tactic struct {
	minAtt int
	maxAtt int
	Name   string
}

func (t *Tactic) Use() int {
	return t.minAtt + rand.Intn(t.maxAtt-t.minAtt)
}

func init() {
	rand.Seed(time.Now().UTC().UnixNano())
}
