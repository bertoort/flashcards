package views

import (
	"github.com/gizak/termui"

	"github.com/berto/flashcards/components"
	"github.com/berto/flashcards/config"
)

const optionsHeight = 3

// View holds the application components
type View struct {
	Flashcard *components.Flashcard
	Options   *components.Options
	WordList  *components.WordList
}

// CreateFlashcardView creates a new application view
func CreateFlashcardView(config *config.Config, index int) *View {
	flashcard := components.CreateFlashcard(&config.Flashcards[index], optionsHeight)
	options := components.CreateOptions(optionsHeight)
	wordList := components.CreateWordList(&config.Flashcards, optionsHeight)

	view := &View{
		Options:   options,
		Flashcard: flashcard,
		WordList:  wordList,
	}

	return view
}

// Refresh re-renderse the components
func (v *View) Refresh() {
	termui.Render(
		v.Flashcard,
		v.WordList,
		v.Options,
	)
}
