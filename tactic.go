package main

import (
	"math/rand"
	"time"
)

type Tactic struct {
	base  int
	bonus int
	Name  string
}

func (t *Tactic) Use() int {
	return t.base + rand.Intn(t.bonus)
}

func init() {
	rand.Seed(time.Now().UTC().UnixNano())
}
