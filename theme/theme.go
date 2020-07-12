// SPDX-License-Identifier: Unlicense OR MIT

package theme

import (
	"gioui.org/font/gofont"
	"gioui.org/text"
	"gioui.org/unit"
	"gioui.org/widget"
	"gioui.org/widget/material"
	"image/color"
)

type DuoUItheme struct {
	T        *material.Theme
	Shaper   text.Shaper
	TextSize unit.Value
	Color    struct {
		Primary color.RGBA
		Text    color.RGBA
		Hint    color.RGBA
		InvText color.RGBA
	}
	Colors map[string]string
	Fonts  map[string]text.Typeface
	Icons  map[string]*widget.Icon

	scrollBarSize int
}

func register(fnt text.Font, ttf []byte) {
	//face, err := opentype.Parse(ttf)
	//if err != nil {
	//	panic(fmt.Sprintf("failed to parse font: %v", err))
	//}
	//fnt.Typeface = "Go"
	//font.Register(fnt, face)
}

func NewDuoUItheme() *DuoUItheme {
	t := &DuoUItheme{
		T: material.NewTheme(gofont.Collection()),
		//Shaper: font.Default(),
	}
	t.Colors = NewDuoUIcolors()
	t.TextSize = unit.Sp(16)
	//t.Icons = NewDuoUIicons()
	return t
}

func (t *DuoUItheme) ChangeLightDark() {
	light := t.Colors["Light"]
	dark := t.Colors["Dark"]
	lightGray := t.Colors["LightGrayIII"]
	darkGray := t.Colors["DarkGrayII"]
	t.Colors["Light"] = dark
	t.Colors["Dark"] = light
	t.Colors["LightGrayIII"] = darkGray
	t.Colors["DarkGrayII"] = lightGray
}
