package panel

import (
	"gioui.org/io/pointer"
	"gioui.org/layout"
	"gioui.org/unit"
	"gioui.org/widget"
	"gioui.org/widget/material"
	"github.com/gioapp/gel/container"
	"github.com/gioapp/gel/helper"
	"github.com/gioapp/gel/theme"
	"image"
)

type DuoUIpanelStyle struct {
	PanelObject func(gtx layout.Context) layout.Dimensions
	ScrollBar   *ScrollBarStyle
	container   container.DuoUIcontainerStyle
}

func DuoUIpanelSt(t *theme.DuoUItheme) DuoUIpanelStyle {
	return DuoUIpanelStyle{
		container: container.DuoUIcontainer(t, 0, t.Colors["Light"]),
	}
}
func (p *DuoUIpanelStyle) panelLayout(panel *Panel, row func(i int, in func(gtx layout.Context) layout.Dimensions)) func(gtx layout.Context) layout.Dimensions {
	return func(gtx layout.Context) layout.Dimensions {
		visibleObjectsNumber := 0
		return panel.PanelContentLayout.Layout(gtx, panel.PanelObjectsNumber, func(gtx layout.Context, i int) layout.Dimensions {
			row(i, p.PanelObject)
			visibleObjectsNumber = visibleObjectsNumber + 1
			panel.VisibleObjectsNumber = visibleObjectsNumber
			return layout.Dimensions{}
		})
	}
}

func (p *DuoUIpanelStyle) Layout(gtx layout.Context, panel *Panel, row func(i int, in func(gtx layout.Context) layout.Dimensions)) layout.Dimensions {
	return p.container.Layout(gtx, layout.NW, func(gtx layout.Context) layout.Dimensions {
		return layout.Flex{
			Axis:    layout.Horizontal,
			Spacing: layout.SpaceBetween,
		}.Layout(gtx,
			layout.Flexed(1, p.panelLayout(panel, row)),
			layout.Rigid(func(gtx layout.Context) layout.Dimensions {
				if panel.PanelObjectsNumber > panel.VisibleObjectsNumber {
					return p.ScrollBarLayout(gtx, panel)
				} else {
					return p.ScrollBarLayout(gtx, panel)
				}
			}),
		)
		//fmt.Println("scrollUnit:", panel.ScrollUnit)
		//fmt.Println("ScrollBar.Slider.Height:", panel.ScrollBar.Slider.Height)
		//fmt.Println("PanelObjectsNumber:", panel.PanelObjectsNumber)

		panel.Layout(&gtx)
		return layout.Dimensions{}
	})
}

var (
	widgetButtonUp   = new(widget.Clickable)
	widgetButtonDown = new(widget.Clickable)
)

type ScrollBarStyle struct {
	ColorBg      string
	BorderRadius [4]float32
	OperateValue interface{}
	slider       *ScrollBarSlider
	up           material.IconButtonStyle
	down         material.IconButtonStyle
	container    container.DuoUIcontainerStyle
}

type ScrollBarSlider struct {
	container container.DuoUIcontainerStyle
	Icon      widget.Icon
}

func ScrollBarSt(t *theme.DuoUItheme, leftMargin int) *ScrollBarStyle {
	slider := &ScrollBarSlider{
		container: container.DuoUIcontainer(t, 0, t.Colors["Primary"]),
		//Icon:      *t.Icons["Grab"],
	}
	slider.container.CornerRadius = 8
	scrollbar := &ScrollBarStyle{
		ColorBg:      t.Colors["DarkGrayII"],
		BorderRadius: [4]float32{},
		slider:       slider,
		//up:           t.IconButton(t.Icons["Up"]),
		//down:         t.IconButton(t.Icons["Down"]),
		container: container.DuoUIcontainer(t, 0, t.Colors["Light"]),
	}
	scrollbar.container.PaddingLeft = leftMargin
	return scrollbar
}

func (p *DuoUIpanelStyle) ScrollBarLayout(gtx layout.Context, panel *Panel) layout.Dimensions {
	return p.ScrollBar.container.Layout(gtx, layout.Center, func(gtx layout.Context) layout.Dimensions {
		return layout.Flex{
			Axis: layout.Vertical,
		}.Layout(gtx,
			layout.Rigid(func(gtx layout.Context) layout.Dimensions {
				for widgetButtonUp.Clicked() {
					if panel.PanelContentLayout.Position.First > 0 {
						panel.PanelContentLayout.Position.First = panel.PanelContentLayout.Position.First - 1
						panel.PanelContentLayout.Position.Offset = 0
					}
				}
				//p.ScrollBar.up.Padding = unit.Dp(0)
				//p.ScrollBar.up.Size = unit.Dp(16)
				p.ScrollBar.up.Color = helper.HexARGB("ffcfcfcf")
				return p.ScrollBar.up.Layout(gtx)
			}),
			layout.Flexed(1, func(gtx layout.Context) layout.Dimensions {
				return p.bodyLayout(gtx, panel)
			}),
			layout.Rigid(func(gtx layout.Context) layout.Dimensions {
				for widgetButtonDown.Clicked() {
					if panel.PanelContentLayout.Position.BeforeEnd {
						panel.PanelContentLayout.Position.First = panel.PanelContentLayout.Position.First + 1
						panel.PanelContentLayout.Position.Offset = 0
					}
				}
				//p.ScrollBar.down.Padding = unit.Dp(0)
				p.ScrollBar.down.Size = unit.Dp(16)
				p.ScrollBar.down.Color = helper.HexARGB("ffcfcfcf")
				return p.ScrollBar.down.Layout(gtx)
			}),
		)
	})
}

func (p *DuoUIpanelStyle) bodyLayout(gtx layout.Context, panel *Panel) layout.Dimensions {
	return layout.Flex{
		Axis: layout.Vertical,
	}.Layout(gtx,
		layout.Rigid(func(gtx layout.Context) layout.Dimensions {
			cs := gtx.Constraints
			pointer.Rect(
				image.Rectangle{Max: image.Point{X: cs.Max.X, Y: cs.Max.Y}},
			).Add(gtx.Ops)
			//pointer.InputOp{Key: panel.ScrollBar.Slider}.Add(gtx.Ops)
			return layout.Center.Layout(gtx, func(gtx layout.Context) layout.Dimensions {
				return layout.Inset{
					Top: unit.Dp(float32(panel.PanelContentLayout.Position.First) * panel.ScrollUnit),
				}.Layout(gtx, func(gtx layout.Context) layout.Dimensions {
					gtx.Constraints.Min.X = panel.ScrollBar.Size
					gtx.Constraints.Min.Y = panel.ScrollBar.Slider.CursorHeight
					if panel.ScrollBar.Slider.CursorHeight < panel.ScrollBar.Size {
						panel.ScrollBar.Slider.CursorHeight = panel.ScrollBar.Size
					}
					return p.ScrollBar.slider.container.Layout(gtx, layout.W, func(gtx layout.Context) layout.Dimensions {
						p.ScrollBar.slider.Icon.Color = helper.HexARGB("ffcfcfcf")
						return p.ScrollBar.slider.Icon.Layout(gtx, unit.Px(float32(panel.ScrollBar.Size)))
					})
				})
				panel.ScrollBar.Slider.Layout(&gtx)
				return layout.Dimensions{}
			})
		}),
	)
}
