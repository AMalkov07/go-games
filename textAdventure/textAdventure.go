package main

import (
	"bufio"
	"fmt"
	"os"
)

type storyPage struct {
	text    string
	options map[string]*storyPage
}

func (sp *storyPage) addOption(command string, text string) {
	newPage := storyPage{text, make(map[string]*storyPage)}
	sp.options[command] = &newPage
}

func (sp *storyPage) printPage() {
	fmt.Println(sp.text)
}

func (sp *storyPage) userInput(scanner *bufio.Scanner) string {
	scanner.Scan()
	return scanner.Text()
}

func (sp *storyPage) playGame() {
	scanner := bufio.NewScanner(os.Stdin)
	var ui string
	for len(sp.options) > 0 {
		sp.printPage()
		ui = sp.userInput(scanner)
		if val, ok := sp.options[ui]; !ok {
			fmt.Println("invalid input")
		} else {
			sp = val
		}
	}
	sp.printPage()
	fmt.Println("end of adventure")
}

func main() {
	firstPage := storyPage{"you wake up in the middle of a room with no memories. To the left is a door, to the righ is a cluttered desk, in front of you is a window. What would you like to do?", make(map[string]*storyPage)}
	firstPage.addOption("go left", "you walk over to the door, the door is locked")
	trav := firstPage.options["go left"]
	trav.addOption("bang on door", "nothing happens")
	firstPage.addOption("go right", "you walk over to the desk, it is cluttered with many objects")
	trav = firstPage.options["go right"]
	trav.addOption("search desk", "the desk contains a knife, a stack of papers, and a flashlight")
	firstPage.addOption("go forward", "you walk over to the window in front of you, you see a forest")
	trav = firstPage.options["go forward"]
	trav.addOption("open window", "the window is locked")
	firstPage.playGame()
}
