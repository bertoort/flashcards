package context

import (
	"log"
	"math/rand"
	"time"

	"github.com/gizak/termui"

	"github.com/berto/flashcards/config"
	"github.com/berto/flashcards/views"
)

type AppContext struct {
	Body         *termui.Grid
	View         *views.View
	Config       *config.Config
	CurrentIndex int
}

// CreateAppContext creates an application context which can be passed
// and referenced througout the application
func CreateAppContext(flgConfig *string, random *bool) *AppContext {
	// Load config
	config, err := config.NewConfig(flgConfig)
	if err != nil {
		log.Fatalf("ERROR: not able to load config file (%s): %s", &flgConfig, err)
	}

	// Find starting index
	var index int
	if *random {
		index = randomIndex(len(config.Flashcards))
	}

	// Create FlashcardView
	view := views.CreateFlashcardView(config, index)

	return &AppContext{
		View:         view,
		Config:       config,
		CurrentIndex: index,
	}
}

func randomIndex(length int) int {
	return rand.New(rand.NewSource(time.Now().UnixNano())).Intn(length - 1)
}
