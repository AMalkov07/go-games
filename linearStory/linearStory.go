package main

import "fmt"

type storyPage struct {
	text     string
	nextPage *storyPage
}

func (sp *storyPage) readStory() {
	if sp == nil {
		return
	}
	fmt.Println(sp.text)
	sp.nextPage.readStory()
}

// we make sure to operate on the pointer of storyPatge instead of just storyPage so that we have pass by reference instead of pass by value
func (sp *storyPage) addPage(s string) {
	newPage := new(storyPage)
	newPage.nextPage = sp.nextPage
	newPage.text = s
	sp.nextPage = newPage
}

func (sp *storyPage) appendPage(s string) {
	for sp.nextPage != nil {
		sp = sp.nextPage
	}
	newPage := new(storyPage)
	newPage.text = s
	sp.nextPage = newPage
}

func main() {
	firstPage := storyPage{"you awaken in an empty house", nil}
	firstPage.appendPage("to the left of you is a table, to the right of you is a door")
	firstPage.appendPage("what would you like to do?")
	firstPage.readStory()
}
