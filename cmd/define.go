package commands

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/berto/flashcards/config"
	"github.com/sajari/fuzzy"
)

func Define(configPath, word *string) {
	conf, err := config.NewConfig(configPath)
	if err != nil {
		log.Fatalf("ERROR: not able to load config file (%s): %s", &configPath, err)
	}

	var found bool
	for i := 0; i < len(conf.Flashcards); i++ {
		if strings.ToLower(conf.Flashcards[i].Name) == strings.ToLower(*word) {
			fmt.Printf("Definition: %s", conf.Flashcards[i].Definition)
			found = true
		}
	}

	if !found {
		model := fuzzy.NewModel()
		model.SetDepth(5)
		model.Train(getWords(conf.Flashcards))
		var suggestions = model.Suggestions(*word, true)
		if len(suggestions) > 0 {
			fmt.Printf("Word not found. Maybe you meant: %v", suggestions)
		} else {
			fmt.Println("Word not found")
		}
	}

	os.Exit(0)
}

func getWords(flashcards []config.Flashcard) (words []string) {
	for i := 0; i < len(flashcards); i++ {
		words = append(words, strings.ToLower(flashcards[i].Name))
	}
	return
}
