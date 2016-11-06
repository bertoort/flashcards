package handlers

import (
	"strconv"

	"github.com/gizak/termui"
	termbox "github.com/nsf/termbox-go"

	"github.com/berto/flashcards/context"
)

// actionMap binds specific action names to the function counterparts
var actionMap = map[string]func(*context.AppContext){
	"q": actionQuit,
	"f": actionFlip,
	// "channel-up":     actionMoveCursorUpChannels,
	// "channel-down":   actionMoveCursorDownChannels,
	// "channel-top":    actionMoveCursorTopChannels,
	// "channel-bottom": actionMoveCursorBottomChannels,
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

// TODO: resize only seems to work for width and resizing it too small
// will cause termui to panic
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
	termui.Render(ctx.View.Flashcard.Par)
}

// func actionGetMessages(ctx *context.AppContext) {
//   ctx.View.Chat.GetMessages(
//     ctx.Service,
//     ctx.Service.Channels[ctx.View.Channels.SelectedChannel],
//   )
//
//   termui.Render(ctx.View.Chat)
// }
//
// func actionMoveCursorUpWords(ctx *context.AppContext) {
//   ctx.View.Channels.MoveCursorUp()
//   termui.Render(ctx.View.Channels)
//   actionChangeChannel(ctx)
// }
//
// func actionMoveCursorDownChannels(ctx *context.AppContext) {
//   go func() {
//     if timer != nil {
//       timer.Stop()
//     }
//
//     ctx.View.Channels.MoveCursorDown()
//     termui.Render(ctx.View.Channels)
//
//     timer = time.NewTimer(time.Second / 4)
//     <-timer.C
//
//     actionChangeChannel(ctx)
//   }()
// }
//
// func actionMoveCursorTopChannels(ctx *context.AppContext) {
//   ctx.View.Channels.MoveCursorTop()
//   actionChangeChannel(ctx)
// }
//
// func actionMoveCursorBottomChannels(ctx *context.AppContext) {
//   ctx.View.Channels.MoveCursorBottom()
//   actionChangeChannel(ctx)
// }
//
// func actionChangeChannel(ctx *context.AppContext) {
//   // Clear messages from Chat pane
//   ctx.View.Chat.ClearMessages()
//
//   // Get message for the new channel
//   ctx.View.Chat.GetMessages(
//     ctx.Service,
//     ctx.Service.SlackChannels[ctx.View.Channels.SelectedChannel],
//   )
//
//   // Set channel name for the Chat pane
//   ctx.View.Chat.SetBorderLabel(
//     ctx.Service.Channels[ctx.View.Channels.SelectedChannel],
//   )
//
//   // Set read mark
//   ctx.View.Channels.SetReadMark(ctx.Service)
//
//   termui.Render(ctx.View.Channels)
//   termui.Render(ctx.View.Chat)
// }
//
// func actionNewMessage(ctx *context.AppContext, channelID string) {
//   ctx.View.Channels.NewMessage(ctx.Service, channelID)
//   termui.Render(ctx.View.Channels)
// }
//

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
