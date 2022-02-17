# Wikiventure

This is a silly spoofish game about mediawiki code review. It is not intended to be taken seriously.

You play as a new developer trying to get your first change to mediawiki merged. Your odds of achieving this are ... well ... you'll see.

## Quickstart:
1) install go (the language, see [go.dev](https://go.dev/))
2) clone this repository to your GOPATH and change to its directory in your terminal
3) execute "go get github.com/fatih/color"
4) execute "go run *.go"

## Notes
You can, of course, execute the game in other ways, including building and running a binary. 

Color schemes are quite rudimentary. If the colored text output doesn't look good on your terminal, type "color none" (in the game) and colored output will be disabled. Preferences are not saved anywhere, so you'll have to do this every time.

This game was inspired by a comment made by [Cindy Cicalese](https://meta.wikimedia.org/wiki/User:CCicalese_(WMF)) during a meeting. Initial implementation was by [Bill Pirkle](https://meta.wikimedia.org/wiki/User:BPirkle_(WMF))

Implementation ideas were taken from [Paul Fortin's blog](https://gocodecloud.com/blog/2016/03/19/writing-a-text-adventure-game-in-go---part-1/). Some of the code is still directly cut-and-paste from there.

Consensus during code review can go up even when the player and the reviewer are being disagreeable to each other, because they are effectively beating each other into submission. This is, of course, not how we want real-life code reviews to go. But this is a spoof. 

## Disclaimer
This is my first-ever Go project. I wrote it to practice with the language. The quality of the code (or lack thereof) probably reflects my level of (un)familiarity with Go in ways that I'm not even yet aware of. 