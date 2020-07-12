package icontextbtn

import (
	"gioui.org/layout"
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
	Word         string
	CornerRadius unit.Value
}

func IconTextBtn(t *material.Theme, b *widget.Clickable, i *widget.Icon, c, w string) IconTextButton {
	return IconTextButton{
		Theme:     t,
		Button:    b,
		Icon:      i,
		IconColor: c,
		Word:      w,
	}
}

func (b IconTextButton) Layout(gtx layout.Context) layout.Dimensions {
	bb := material.ButtonLayout(b.Theme, b.Button)
	bb.CornerRadius = b.CornerRadius
	bb.Background = b.Background

	return bb.Layout(gtx, func(gtx layout.Context) layout.Dimensions {
		return layout.UniformInset(unit.Dp(8)).Layout(gtx, func(gtx layout.Context) layout.Dimensions {
			gtx.Constraints.Min.X = gtx.Constraints.Max.X
			iconAndLabel := layout.Flex{Axis: layout.Horizontal, Alignment: layout.Middle, Spacing: layout.SpaceBetween}
			textIconSpacer := unit.Dp(8)

			layIcon := layout.Rigid(func(gtx layout.Context) layout.Dimensions {
				return layout.Inset{Right: textIconSpacer}.Layout(gtx, func(gtx layout.Context) layout.Dimensions {
					var d layout.Dimensions
					if b.Icon != nil {
						size := gtx.Px(unit.Dp(56)) - 2*gtx.Px(unit.Dp(16))
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
				return layout.Inset{Left: textIconSpacer}.Layout(gtx, func(gtx layout.Context) layout.Dimensions {
					l := material.Body1(b.Theme, b.Word)
					l.Color = b.Theme.Color.InvText
					return l.Layout(gtx)
				})
			})

			return iconAndLabel.Layout(gtx, layIcon, layLabel)
		})
	})
}
