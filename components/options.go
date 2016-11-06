package components

import "github.com/gizak/termui"

// Options is the definition of Options component
type Options struct {
	Par *termui.Par
}

// CreateOptions is the constructor for the Options component
func CreateOptions(height int) *Options {
	options := []string{
		"[q]quit",
		"[f]flip",
		"[k]up",
		"[j]down",
		"[s]select",
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

// Buffer implements interface termui.Bufferer
func (o *Options) Buffer() termui.Buffer {
	buf := o.Par.Buffer()
	return buf
}

// GetHeight implements interface termui.GridBufferer
func (o *Options) GetHeight() int {
	return o.Par.Block.GetHeight()
}

// SetWidth implements interface termui.GridBufferer
func (o *Options) SetWidth(w int) {
	o.Par.SetWidth(w)
}

// SetX implements interface termui.GridBufferer
func (o *Options) SetX(x int) {
	o.Par.SetX(x)
}

// SetY implements interface termui.GridBufferer
func (o *Options) SetY(y int) {
	o.Par.SetY(y)
}
