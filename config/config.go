package config

import (
	"encoding/json"
	"errors"
	"os"
)

// Config is the definition of a Config struct
type Config struct {
	Flashcards []Flashcard `json:"flashcards"`
}

// Flashard is the definition of a flashcard
type Flashcard struct {
	Name       string `json:"name"`
	Definition string `json:"definition"`
}

// NewConfig loads the config file and returns a Config struct
func NewConfig(filepath *string) (cfg *Config, err error) {
	file, err := os.Open(*filepath)
	if err != nil {
		return cfg, err
	}

	if err := json.NewDecoder(file).Decode(&cfg); err != nil {
		return cfg, err
	}

	if len(cfg.Flashcards) == 0 {
		return cfg, errors.New("No flashcards found")
	}

	return cfg, nil
}
