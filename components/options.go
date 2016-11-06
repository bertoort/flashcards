package components

import "github.com/gizak/termui"

type Options struct {
	Par *termui.Par
}

func CreateOptions(height int) *Options {
	options := []string{
		"[f]flip",
		"[n]next",
		"[p]previous",
		"[r]random",
		"[q]quit",
	}
	var text string
	spacing := " "
	for i := 0; i < len(options); i++ {
		text += options[i] + spacing
	}
	par := termui.NewPar(text)
	par.BorderLabel = "Options"
	par.Height = height
	return &Options{
		Par: par,
	}
}
