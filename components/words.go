package components

import (
	"github.com/berto/flashcards/config"
	"github.com/gizak/termui"
)

type WordList struct {
	List *termui.List
}

func CreateWordList(flashcards *[]config.Flashcard, optionsHeight int) *WordList {
	wordList := getNames(*flashcards)
	list := termui.NewList()
	list.Items = wordList
	list.ItemFgColor = termui.ColorYellow
	list.BorderLabel = "Words"
	list.Height = 43
	list.Height = termui.TermHeight() - optionsHeight
	return &WordList{
		List: list,
	}
}

func getNames(wordList []config.Flashcard) []string {
	words := make([]string, len(wordList))
	for i := 0; i < len(wordList); i++ {
		words = append(words, wordList[i].Name)
	}
	return words
}
