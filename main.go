package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"time"
)

var Out *os.File
var In *os.File

func init() {
	rand.Seed(time.Now().UTC().UnixNano())
	Out = os.Stdout
	In = os.Stdin
}

func main() {
	var game = *new(Game)
	game.Play()
}

func Outputf(color string, format string, args ...interface{}) {
	s := fmt.Sprintf(format, args...)
	Output(color, s)
}

func Output(color string, args ...interface{}) {
	s := fmt.Sprint(args...)

	// The code I copied this from used a package that colored the output
	// It gave me an error, so I removed it for now, but left the color
	// parameter because that sounds cool and I'd like to add it back.

	fmt.Fprintln(Out, s)
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
