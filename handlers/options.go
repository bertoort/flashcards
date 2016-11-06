package handlers

import (
	"strconv"

	"github.com/gizak/termui"
	termbox "github.com/nsf/termbox-go"

	"github.com/berto/flashcards/context"
)

// actionMap binds specific action names to the function counterparts
var actionMap = map[string]func(*context.AppContext){
	"q":        actionQuit,
	"<escape>": actionQuit,
	"f":        actionFlip,
	"k":        actionScrollUp,
	"j":        actionScrollDown,
	"s":        actionSelect,
	"<enter>":  actionSelect,
}

func RegisterEventHandlers(ctx *context.AppContext) {
	anyKeyHandler(ctx)
	termui.Handle("/sys/wnd/resize", resizeHandler(ctx))
}

func anyKeyHandler(ctx *context.AppContext) {
	go func() {
		for {
			ev := termbox.PollEvent()
			if ev.Type != termbox.EventKey {
				continue
			}
			keyStr := getKeyString(ev)
			action, ok := actionMap[keyStr]
			if ok {
				action(ctx)
			}
		}
	}()
}

func resizeHandler(ctx *context.AppContext) func(termui.Event) {
	return func(e termui.Event) {
		actionResize(ctx)
	}
}

func actionResize(ctx *context.AppContext) {
	termui.Body.Width = termui.TermWidth()
	termui.Body.Align()
	termui.Render(termui.Body)
}

func actionQuit(*context.AppContext) {
	termui.StopLoop()
}

func actionFlip(ctx *context.AppContext) {
	ctx.View.Flashcard.Flip()
	termui.Render(ctx.View.Flashcard)
}

func actionScrollUp(ctx *context.AppContext) {
	change := ctx.View.WordList.ScrollUp()
	if change {
		termui.Render(ctx.View.WordList)
	}
}

func actionScrollDown(ctx *context.AppContext) {
	change := ctx.View.WordList.ScrollDown()
	if change {
		termui.Render(ctx.View.WordList)
	}
}

func actionSelect(ctx *context.AppContext) {
	ctx.View.UpdateFlashcard()
	termui.Render(ctx.View.Flashcard)
}

// See:
// - https://github.com/gizak/termui/blob/a7e3aeef4cdf9fa2edb723b1541cb69b7bb089ea/events.go#L31-L72
// - https://github.com/nsf/termbox-go/blob/master/api_common.go
func getKeyString(e termbox.Event) string {
	var ek string

	k := string(e.Ch)
	pre := ""
	mod := ""

	if e.Mod == termbox.ModAlt {
		mod = "M-"
	}
	if e.Ch == 0 {
		if e.Key > 0xFFFF-12 {
			k = "<f" + strconv.Itoa(0xFFFF-int(e.Key)+1) + ">"
		} else if e.Key > 0xFFFF-25 {
			ks := []string{"<insert>", "<delete>", "<home>", "<end>", "<previous>", "<next>", "<up>", "<down>", "<left>", "<right>"}
			k = ks[0xFFFF-int(e.Key)-12]
		}

		if e.Key <= 0x7F {
			pre = "C-"
			k = string('a' - 1 + int(e.Key))
			kmap := map[termbox.Key][2]string{
				termbox.KeyCtrlSpace:     {"C-", "<space>"},
				termbox.KeyBackspace:     {"", "<backspace>"},
				termbox.KeyTab:           {"", "<tab>"},
				termbox.KeyEnter:         {"", "<enter>"},
				termbox.KeyEsc:           {"", "<escape>"},
				termbox.KeyCtrlBackslash: {"C-", "\\"},
				termbox.KeyCtrlSlash:     {"C-", "/"},
				termbox.KeySpace:         {"", "<space>"},
				termbox.KeyCtrl8:         {"C-", "8"},
			}
			if sk, ok := kmap[e.Key]; ok {
				pre = sk[0]
				k = sk[1]
			}
		}
	}

	ek = pre + mod + k
	return ek
}
