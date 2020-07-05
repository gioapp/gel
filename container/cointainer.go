package container

import (
	"gioui.org/f32"
	"gioui.org/layout"
	"gioui.org/op/clip"
	"gioui.org/text"
	"gioui.org/unit"
	"github.com/gioapp/gel/helper"
	"github.com/gioapp/gel/theme"
)

type DuoUIcontainerStyle struct {
	// Color is the text color.
	Color         string
	Font          text.Font
	TextSize      unit.Value
	Background    string
	TxColor       string
	BgColor       string
	TxColorHover  string
	BgColorHover  string
	FullWidth     bool
	Width         int
	Height        int
	CornerRadius  int
	PaddingTop    int
	PaddingRight  int
	PaddingBottom int
	PaddingLeft   int
	shaper        text.Shaper
	link          bool
	hover         bool
}

func DuoUIcontainer(t *theme.DuoUItheme, padding int, background string) DuoUIcontainerStyle {
	return DuoUIcontainerStyle{
		Font: text.Font{
			Typeface: t.Fonts["Primary"],
		},
		// Color:      rgb(0xffffff),
		PaddingTop:    padding,
		PaddingRight:  padding,
		PaddingBottom: padding,
		PaddingLeft:   padding,
		Background:    background,
		TextSize:      t.TextSize.Scale(14.0 / 16.0),
		shaper:        t.Shaper,
	}
}

func (d DuoUIcontainerStyle) Layout(gtx layout.Context, direction layout.Direction, itemContent func(gtx layout.Context) layout.Dimensions) layout.Dimensions {
	//hmin := gtx.Constraints.Min.X
	//vmin := gtx.Constraints.Min.Y
	//if d.FullWidth {
	//	hmin = gtx.Constraints.Max.Y
	//}
	return layout.Stack{Alignment: layout.W}.Layout(gtx,
		layout.Expanded(func(gtx layout.Context) layout.Dimensions {
			rr := float32(gtx.Px(unit.Dp(float32(d.CornerRadius))))
			clip.Rect{
				Rect: f32.Rectangle{Max: f32.Point{
					X: float32(gtx.Constraints.Min.X),
					Y: float32(gtx.Constraints.Min.Y),
				}},
				NE: rr, NW: rr, SE: rr, SW: rr,
			}.Op(gtx.Ops).Add(gtx.Ops)
			return helper.Fill(gtx, helper.HexARGB(d.Background))
			//return pointer.Rect(image.Rectangle{Max: gtx.Now.Size}).Add(gtx.Ops)
		}),
		layout.Stacked(func(gtx layout.Context) layout.Dimensions {
			//gtx.Constraints.Min = hmin
			//gtx.Constraints.Min = vmin
			return direction.Layout(gtx, func(gtx layout.Context) layout.Dimensions {
				return layout.Inset{
					Top:    unit.Dp(float32(d.PaddingTop)),
					Right:  unit.Dp(float32(d.PaddingRight)),
					Bottom: unit.Dp(float32(d.PaddingBottom)),
					Left:   unit.Dp(float32(d.PaddingLeft)),
				}.Layout(gtx, itemContent)
			})
		}),
	)
}
