package helper

import (
	"gioui.org/f32"
	"gioui.org/layout"
	"gioui.org/op/paint"
	"gioui.org/unit"
	"image"
)

func DuoUIdrawRectangle(gtx layout.Context, w, h int, color string, borderRadius [4]float32, padding [4]float32) layout.Dimensions {
	in := layout.Inset{
		Top:    unit.Dp(padding[0]),
		Right:  unit.Dp(padding[1]),
		Bottom: unit.Dp(padding[2]),
		Left:   unit.Dp(padding[3]),
	}
	return in.Layout(gtx, func(gtx layout.Context) layout.Dimensions {
		square := f32.Rectangle{
			Max: f32.Point{
				X: float32(w),
				Y: float32(h),
			},
		}
		paint.ColorOp{Color: HexARGB(color)}.Add(gtx.Ops)
		//clip.Rect{Rect: square,
		//	NE: borderRadius[0], NW: borderRadius[1], SE: borderRadius[2], SW: borderRadius[3]}.Op(gtx.Ops).Add(gtx.Ops) // HLdraw
		paint.PaintOp{Rect: square}.Add(gtx.Ops)
		return layout.Dimensions{Size: image.Point{X: w, Y: h}}
	})
}

func DuoUIfill(gtx layout.Context, col string) {
	cs := gtx.Constraints
	d := image.Point{X: cs.Min.X, Y: cs.Min.Y}
	dr := f32.Rectangle{
		Max: f32.Point{X: float32(d.X), Y: float32(d.Y)},
	}
	paint.ColorOp{Color: HexARGB(col)}.Add(gtx.Ops)
	paint.PaintOp{Rect: dr}.Add(gtx.Ops)
	//gtx.Dimensions = layout.Dimensions{Size: d}
}

func DuoUIline(vert bool, verticalPadding, horizontalPadding float32, size int, color string) func(gtx layout.Context) layout.Dimensions {
	return func(gtx layout.Context) layout.Dimensions {
		return layout.Inset{
			Top:    unit.Dp(verticalPadding),
			Right:  unit.Dp(horizontalPadding),
			Bottom: unit.Dp(verticalPadding),
			Left:   unit.Dp(horizontalPadding),
		}.Layout(gtx, func(gtx layout.Context) layout.Dimensions {
			v := size
			h := gtx.Constraints.Max.X
			if vert {
				h = size
				v = 8
			}
			return DuoUIdrawRectangle(gtx, h, v, color, [4]float32{0, 0, 0, 0}, [4]float32{0, 0, 0, 0})
		})
	}
}

func toPointF(p image.Point) f32.Point {
	return f32.Point{X: float32(p.X), Y: float32(p.Y)}
}
