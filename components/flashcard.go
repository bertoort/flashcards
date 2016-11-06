package components

import (
	"github.com/berto/flashcards/config"
	"github.com/gizak/termui"
)

// Flashcard is the definition of Flashcard component
type Flashcard struct {
	Par               *termui.Par
	name              string
	definition        string
	displayDefinition bool
	optionsHeight     int
}

// CreateFlashcard is the constructor for the Flashcard component
func CreateFlashcard(flashcard *config.Flashcard, optionsHeight int) *Flashcard {
	return &Flashcard{
		Par:               newFlashcard(flashcard.Name, "Term", optionsHeight),
		name:              flashcard.Name,
		definition:        flashcard.Definition,
		displayDefinition: true,
	}
}

// Flip displays name or definition accordingly
func (f *Flashcard) Flip() {
	f.Par.BorderLabel = "Term"
	f.Par.Text = f.name
	if f.displayDefinition {
		f.Par.BorderLabel = "Definition"
		f.Par.Text = f.definition
	}
	f.displayDefinition = !f.displayDefinition
}

// Buffer implements interface termui.Bufferer
func (f *Flashcard) Buffer() termui.Buffer {
	buf := f.Par.Buffer()
	return buf
}

// GetHeight implements interface termui.GridBufferer
func (f *Flashcard) GetHeight() int {
	return f.Par.Block.GetHeight()
}

// SetWidth implements interface termui.GridBufferer
func (f *Flashcard) SetWidth(w int) {
	f.Par.SetWidth(w)
}

// SetX implements interface termui.GridBufferer
func (f *Flashcard) SetX(x int) {
	f.Par.SetX(x)
}

// SetY implements interface termui.GridBufferer
func (f *Flashcard) SetY(y int) {
	f.Par.SetY(y)
}

func newFlashcard(text, label string, optionsHeight int) *termui.Par {
	par := termui.NewPar(text)
	par.Height = termui.TermHeight() - optionsHeight
	par.BorderLabel = label
	return par
}
