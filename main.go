package main

import (
	"flag"
	"fmt"
	"log"
	"os/user"
	"path"

	"github.com/berto/flashcards/context"
	"github.com/berto/flashcards/handlers"
	"github.com/gizak/termui"
)

const (
	VERSION = "v0.0.1"
	USAGE   = `NAME:
    flashcards - flashcards for your terminal

USAGE:
    flashcards

VERSION:
    %s

GLOBAL OPTIONS:
   --help, -h
`
)

var (
	flgConfig string
	flgUsage  bool
)

func init() {
	// Get home dir for config file default
	usr, err := user.Current()
	if err != nil {
		log.Fatal(err)
	}

	// Parse flags
	flag.StringVar(
		&flgConfig,
		"config",
		path.Join(usr.HomeDir, ".flashcards.json"),
		"location of config file",
	)

	flag.Usage = func() {
		fmt.Printf(USAGE, VERSION)
	}

	flag.Parse()
}

func main() {
	// start terminal user interface
	err := termui.Init()
	if err != nil {
		log.Fatal(err)
	}
	defer termui.Close()

	random := false

	// create context
	ctx := context.CreateAppContext(&flgConfig, &random)

	// setup body
	termui.Body.AddRows(
		termui.NewRow(
			termui.NewCol(2, 0, ctx.View.WordList.List),
			termui.NewCol(10, 0, ctx.View.Flashcard.Par),
		),
		termui.NewRow(
			termui.NewCol(12, 0, ctx.View.Options.Par),
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
