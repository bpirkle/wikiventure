package wikiventure

import (
	"bufio"
	"fmt"
	"github.com/fatih/color"
	"math/rand"
	"os"
	"time"
)

// Game top-level structure for the game
// Not doing much with this yet, but it may eventually prove useful
type Game struct {
	Player      Actor
	ColorScheme string
}

var Out *os.File
var In *os.File

func init() {
	rand.Seed(time.Now().UTC().UnixNano())
	Out = os.Stdout
	In = os.Stdin
}

func (g *Game) Play() {
	g.ColorScheme = "dark"

	// Should we prompt people for their name instead?
	g.Player = *new(Actor)
	g.Player.Name = "Jimmy Wales"
	g.Player.Morale = 100
	g.Player.Actions = []int{1, 2, 3, 4, 5, 6}
	g.Player.CurrentLocation = "CommandLine"

	lastLocation := ""

	Output("blue", Messages["welcome"])
	for {
		if g.Player.CurrentLocation == lastLocation {
			Output("red", "You haven't gone anywhere. Type 'help' for available commands.")
		} else {
			lastLocation = g.Player.CurrentLocation
			Output("blue", LocationMap[g.Player.CurrentLocation].Description)

			// We really shouldn't process an event unless location has changed.
			// Otherwise you can stay in AFK forever and get crazy morale by hitting Return over and over
			// And you can also get a Story event immediately after a CodeReview event without entering
			// a command, which is a little confusing.
			g.ProcessEvents(LocationMap[g.Player.CurrentLocation].Events)
			if g.Player.Morale <= 0 {
				Output("white", "\nYou have given up hope on your change. Game over.")
				return
			} else {
				Output("white", "\tYou are still working on your change.")
			}
			Output("blue", "\tMorale:", g.Player.Morale)
		}

		Output("green", "You can go to these places:")
		for _, loc := range LocationMap[g.Player.CurrentLocation].Transitions {
			Outputf("green", "\t%s", loc)
		}
		cmd := UserInputln()
		ProcessCommands(g, cmd)
	}
}

func (g *Game) ProcessEvents(events []string) {
	for _, evtName := range events {
		g.Player.Morale += Events[evtName].ProcessEvent(&g.Player)
	}
}

func Outputf(c string, format string, args ...interface{}) {
	s := fmt.Sprintf(format, args...)
	Output(c, s)
}

func Output(c string, args ...interface{}) {
	s := fmt.Sprint(args...)

	col := color.WhiteString
	switch c {
	case "green":
		col = color.GreenString
	case "red":
		col = color.RedString
	case "blue":
		col = color.BlueString
	case "yellow":
		col = color.YellowString
	}
	fmt.Fprintln(Out, col(s))

	//	fmt.Fprintln(Out, s)
}

func UserInput(i *int) {
	fmt.Fscan(In, i)
}

func UserInputln() string {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("\n >>> ")
	text, _ := reader.ReadString('\n')
	return text
}

func UserInputContinue() string {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("\n Press return to continue the code review")
	text, _ := reader.ReadString('\n')
	return text
}

func (g *Game) setColorScheme(color string) {
	switch color {
	case "dark":
		g.ColorScheme = "dark"
	case "light":
		g.ColorScheme = "light"
	case "none":
		g.ColorScheme = "none"
	default:
		Output("red", "Unrecognized color scheme.")
	}
}
