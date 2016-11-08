package context

import (
	"log"

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
func CreateAppContext(configPath *string) *AppContext {
	// Load config
	config, err := config.NewConfig(configPath)
	if err != nil {
		log.Fatalf("ERROR: not able to load config file (%s): %s", &configPath, err)
	}

	// Create FlashcardView
	view := views.CreateFlashcardView(config)

	return &AppContext{
		View:   view,
		Config: config,
	}
}
