package main

import (
	"flag"
	"fmt"
	"log"
	"os/user"
	"path"

	"github.com/berto/flashcards/cmd"
	"github.com/berto/flashcards/context"
	"github.com/berto/flashcards/handlers"
	"github.com/gizak/termui"
)

const (
	VERSION = "v0.0.1"
	USAGE   = `Name:
    flashcards - flashcards for your terminal

Usage:
    flashcards [-options] [word]

Version:
    %s

Global Options:
   --help, -h
   --config, -c
`
)

var (
	configPath string
	word       *string
)

func init() {
	// Get home dir for config file default
	usr, err := user.Current()
	if err != nil {
		log.Fatal(err)
	}

	// Parse flags
	flag.StringVar(
		&configPath,
		"config",
		path.Join(usr.HomeDir, ".flashcards.json"),
		"location of config file",
	)

	flag.StringVar(
		&configPath,
		"c",
		path.Join(usr.HomeDir, ".flashcards.json"),
		"shorthand for location of config file",
	)

	flag.Usage = func() {
		fmt.Printf(USAGE, VERSION)
	}

	flag.Parse()

	if len(flag.Args()) > 0 {
		word = &flag.Args()[0]
	}
}

func main() {
	if word != nil {
		commands.Define(&configPath, word)
	}

	// start terminal user interface
	err := termui.Init()
	if err != nil {
		log.Fatal(err)
	}
	defer termui.Close()

	// create context
	ctx := context.CreateAppContext(&configPath)

	// setup body
	termui.Body.AddRows(
		termui.NewRow(
			termui.NewCol(2, 0, ctx.View.WordList),
			termui.NewCol(10, 0, ctx.View.Flashcard),
		),
		termui.NewRow(
			termui.NewCol(12, 0, ctx.View.Options),
		),
	)
	termui.Body.Align()
	termui.Render(termui.Body)

	// set body in context
	ctx.Body = termui.Body

	// register handlers
	handlers.RegisterEventHandlers(ctx)

	termui.Loop()
}
