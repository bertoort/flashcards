package components

import (
	"github.com/berto/flashcards/config"
	"github.com/gizak/termui"
)

type Flashcard struct {
	Par               *termui.Par
	name              string
	definition        string
	displayDefinition bool
	optionsHeight     int
}

func CreateFlashcard(flashcard *config.Flashcard, optionsHeight int) *Flashcard {
	return &Flashcard{
		Par:               newFlashcard(flashcard.Name, "name", optionsHeight),
		name:              flashcard.Name,
		definition:        flashcard.Definition,
		displayDefinition: false,
		optionsHeight:     optionsHeight,
	}
}

func newFlashcard(text, label string, optionsHeight int) *termui.Par {
	par := termui.NewPar(text)
	par.Height = termui.TermHeight() - optionsHeight
	par.BorderLabel = label
	return par
}

func (f *Flashcard) Flip() {
	f.Par = newFlashcard(f.name, "name", f.optionsHeight)
	if f.displayDefinition {
		f.Par = newFlashcard(f.definition, "definition", f.optionsHeight)
	}
	f.displayDefinition = !f.displayDefinition
}
