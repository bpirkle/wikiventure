package main

var Messages = map[string]string{
	"welcome": "Welcome to wikiventure!\nYou may type 'help' for help, or 'quit' to exit.\n",
}

// Leftover from code I copied, not currently using these
// Using items seemed buggy, so if we ever reactive this,
// we may have to fix some things.
var Items = map[int]*Item{
	1: {Name: "Key"},
	2: {Name: "Chest", ItemForUse: 1, Contains: []int{3}},
	3: {Name: "Medal"},
}

var Tactics = map[int]*Tactic{
	1: {Name: "Good suggestion", minAtt: 5, maxAtt: 15},
	2: {Name: "Bad suggestion", minAtt: 1, maxAtt: 15},
	3: {Name: "Thoughtful comment", minAtt: 3, maxAtt: 12},
	4: {Name: "Snarky comment", minAtt: 3, maxAtt: 12},
}

var Reviewers = map[int]*Actor{
	0: {Name: "Helpful Novice", Morale: 50, Tactic: 1, Npc: true},
	1: {Name: "Helpful Peer", Morale: 55, Tactic: 1, Npc: true},
	2: {Name: "Helpful Mentor", Morale: 55, Tactic: 3, Npc: true},
	3: {Name: "Hurtful Novice", Morale: 50, Tactic: 2, Npc: true},
	4: {Name: "Hurtful Peer", Morale: 55, Tactic: 2, Npc: true},
	5: {Name: "Hurtful Mentor", Morale: 55, Tactic: 4, Npc: true},
}

var Events = map[string]*Event{
	"codeReview":            {Type: "CodeReview", Chance: 100, Description: "You receive a Code Review comment.", Morale: 0, Evt: ""},
	"unsolicitedCriticism":  {Type: "Story", Chance: 20, Description: "Out of nowhere, someone criticises your code.", Morale: -50, Evt: ""},
	"surpriseEncouragement": {Type: "Story", Chance: 10, Description: "A colleague sends you wikilove.", Morale: +50, Evt: ""},
	"insightfulComment":     {Type: "Story", Chance: 50, Description: "Someone makes a great suggestion", Morale: +30, Evt: ""},
	"extraHoliday":          {Type: "Story", Chance: 10, Description: "You are granted an extra holiday", Morale: +30, Evt: "recharging"},
	"recharging":            {Type: "Story", Chance: 100, Description: "Doing non-computer things you enjoy improves your morale.", Morale: +10, Evt: ""},
}

var LocationMap = map[string]*Location{
	"CommandLine": {Description: "You just pushed your first change.", Transitions: []string{"Phabricator, Gerrit", "Chat", "Meeting", "AFK"}, Events: []string{}},
	"Phab":        {Description: "You are looking at your Phabricator task.", Transitions: []string{"Gerrit", "Chat", "Meeting", "AFK"}, Events: []string{"codeReview", "unsolicitedCriticism"}},
	"Gerrit":      {Description: "You are looking at  your change in Gerrit.", Transitions: []string{"Phab", "Chat", "Meeting", "AFK"}, Events: []string{}}, // Items: []int{2}},
	"Chat":        {Description: "You are in Slack/Element/IRC/whatever-you-prefer.", Transitions: []string{"Phab", "Gerrit", "Meeting", "AFK"}, Events: []string{"insightfulComment"}},
	"Meeting":     {Description: "You are in Google Meet", Transitions: []string{"Phab", "Gerrit", "Chat", "AFK"}, Events: []string{"unsolicitedCriticism"}}, // Items: []int{1}},
	"AFK":         {Description: "You are away from your computer, living your real life", Transitions: []string{"Phab", "Gerrit", "Chat", "Meeting"}, Events: []string{"recharging"}},
}
