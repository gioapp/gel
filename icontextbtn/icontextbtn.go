package icontextbtn

import (
	"gioui.org/layout"
	"gioui.org/text"
	"gioui.org/unit"
	"gioui.org/widget"
	"gioui.org/widget/material"
	"github.com/gioapp/gel/helper"
	"image"
	"image/color"
)

type IconTextButton struct {
	Theme        *material.Theme
	Button       *widget.Clickable
	Background   color.RGBA
	Icon         *widget.Icon
	IconColor    string
	IconSize     unit.Value
	Text         string
	TextSize     unit.Value
	CornerRadius unit.Value
	Axis         layout.Axis
}

func IconTextBtn(t *material.Theme, b *widget.Clickable, i *widget.Icon, is unit.Value, c, w string) IconTextButton {
	return IconTextButton{
		Theme:     t,
		Button:    b,
		Icon:      i,
		IconColor: c,
		IconSize:  is,
		Text:      w,
		TextSize:  unit.Dp(16),
	}
}

func (b IconTextButton) Layout(gtx layout.Context) layout.Dimensions {
	bb := material.ButtonLayout(b.Theme, b.Button)
	bb.CornerRadius = b.CornerRadius
	bb.Background = b.Background
	return bb.Layout(gtx, func(gtx layout.Context) layout.Dimensions {
		return layout.UniformInset(unit.Dp(4)).Layout(gtx, func(gtx layout.Context) layout.Dimensions {
			gtx.Constraints.Min.X = gtx.Constraints.Max.X
			var razmak float32
			if b.Axis != layout.Horizontal {
				razmak = 4
			}
			iconAndLabel := layout.Flex{Axis: b.Axis, Alignment: layout.Middle, Spacing: layout.SpaceBetween}
			layIcon := layout.Rigid(func(gtx layout.Context) layout.Dimensions {
				return layout.Inset{}.Layout(gtx, func(gtx layout.Context) layout.Dimensions {
					var d layout.Dimensions
					if b.Icon != nil {
						//size := gtx.Px(b.IconSize) - 2*gtx.Px(unit.Dp(16))
						size := gtx.Px(b.IconSize) - 2*gtx.Px(unit.Dp(16))
						b.Icon.Color = helper.HexARGB(b.IconColor)
						b.Icon.Layout(gtx, unit.Px(float32(size)))
						d = layout.Dimensions{
							Size: image.Point{X: size, Y: size},
						}
					}
					return d
				})
			})

			layLabel := layout.Rigid(func(gtx layout.Context) layout.Dimensions {
				return layout.Inset{Top: unit.Dp(razmak)}.Layout(gtx, func(gtx layout.Context) layout.Dimensions {
					l := material.Body1(b.Theme, b.Text)
					l.TextSize = unit.Dp(12)
					l.Alignment = text.Middle
					l.Color = b.Theme.Color.InvText
					return l.Layout(gtx)
				})
			})
			layOne := layIcon
			layTwo := layLabel
			if b.Axis != layout.Vertical {
				layOne = layLabel
				layTwo = layIcon
			}
			return iconAndLabel.Layout(gtx, layOne, layTwo)
		})
	})
}
