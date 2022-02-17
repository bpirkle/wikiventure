package wikiventure

var Messages = map[string]string{
	"welcome": "Welcome to wikiventure!\nYou are a new developer trying to get your first change merged\nThe game ends when you either get your patch merged, or run out of Morale and abandon it.\nType 'help' for help, or 'quit' to exit.\n",
}

var Actions = map[int]*Action{
	1: {Name: "Good suggestion", base: 1, bonus: 15},
	2: {Name: "Great suggestion", base: 10, bonus: 15},
	3: {Name: "Bad suggestion", base: -15, bonus: 15},
	4: {Name: "Terrible suggestion", base: -25, bonus: 15},
	5: {Name: "Thoughtful comment", base: 3, bonus: 12},
	6: {Name: "Snarky comment", base: -14, bonus: 12},
}

var Reviewers = map[int]*Actor{
	0: {Name: "Helpful Novice", Morale: 50, Actions: []int{1, 3, 4, 5}, Npc: true},
	1: {Name: "Helpful Peer", Morale: 55, Actions: []int{1, 2, 3, 5}, Npc: true},
	2: {Name: "Helpful Mentor", Morale: 55, Actions: []int{1, 2, 5}, Npc: true},
	3: {Name: "Hurtful Novice", Morale: 50, Actions: []int{1, 3, 4, 6}, Npc: true},
	4: {Name: "Hurtful Peer", Morale: 55, Actions: []int{1, 3, 4, 6}, Npc: true},
	5: {Name: "Hurtful Mentor", Morale: 55, Actions: []int{1, 3, 4, 6}, Npc: true},
}

var Events = map[string]*Event{
	"codeReview":   {Type: "CodeReview", Chance: 25, Description: "You receive a Code Review comment", Morale: 0, Evt: ""},
	"criticism":    {Type: "Story", Chance: 25, Description: "Someone criticises you unfairly", Morale: -40, Evt: ""},
	"wikilove":     {Type: "Story", Chance: 25, Description: "A colleague sends you wikilove", Morale: +40, Evt: ""},
	"unbreakNow":   {Type: "Story", Chance: 25, Description: "You broke the wikis", Morale: -30, Evt: ""},
	"extraHoliday": {Type: "Story", Chance: 20, Description: "You are granted an extra holiday", Morale: +30, Evt: "recharging"},
	"recharging":   {Type: "Story", Chance: 100, Description: "Doing non-computer things you enjoy improves your morale", Morale: +20, Evt: ""},
}

var LocationMap = map[string]*Location{
	"CommandLine": {Description: "You just pushed your first change.", Transitions: []string{"Phab", "Gerrit", "Gitlab", "Chat", "Meeting", "AFK"}, Events: []string{}},
	"Phab":        {Description: "You are looking at your Phabricator task.", Transitions: []string{"Gerrit", "Gitlab", "Chat", "Meeting", "AFK"}, Events: []string{"codeReview", "criticism", "wikilove", "unbreakNow"}},
	"Gerrit":      {Description: "You are looking at  your change in Gerrit.", Transitions: []string{"Phab", "Gitlab", "Chat", "Meeting", "AFK"}, Events: []string{"codeReview", "criticism", "wikilove", "unbreakNow"}},
	"Gitlab":      {Description: "You are in Gitlab. There aren't many people here.", Transitions: []string{"Phab", "Gerrit", "Chat", "Meeting", "AFK"}, Events: []string{"codeReview", "criticism", "wikilove", "unbreakNow"}},
	"Chat":        {Description: "You are in Slack/Element/IRC/whatever-you-prefer.", Transitions: []string{"Phab", "Gerrit", "Gitlab", "Meeting", "AFK"}, Events: []string{"codeReview", "criticism", "wikilove", "unbreakNow"}},
	"Meeting":     {Description: "You are in Google Meet", Transitions: []string{"Phab", "Gerrit", "Gitlab", "Chat"}, Events: []string{"codeReview", "criticism", "unbreakNow", "wikilove", "extraHoliday"}},
	"AFK":         {Description: "You are away from your computer, living your real life", Transitions: []string{"Phab", "Gerrit", "Gitlab", "Chat", "Meeting"}, Events: []string{"recharging"}},
}
