package page

import (
	"gioui.org/layout"
	"gioui.org/text"
	"gioui.org/unit"
	"github.com/gioapp/gel/theme"
)

type DuoUIpage struct {
	Name string
	//Command func()
	Data interface{}

	Title   string
	TxColor string
	// Font          text.Font
	shaper  text.Shaper
	Command func(gtx layout.Context) layout.Dimensions

	Header        func(gtx layout.Context) layout.Dimensions
	HeaderBgColor string
	HeaderPadding float32
	// header
	// header
	Border      float32
	BorderColor string
	Body        func(gtx layout.Context) layout.Dimensions
	BodyBgColor string
	BodyPadding float32
	// body
	// body
	Footer        func(gtx layout.Context) layout.Dimensions
	FooterBgColor string
	FooterPadding float32
	// footer
	// footer
}

type DuoUIpages *map[string]*DuoUIpage

func NewDuoUIpage(t *theme.DuoUItheme, p DuoUIpage) *DuoUIpage {
	return &DuoUIpage{
		Title: p.Title,
		// Font: text.Font{
		// Size: t.TextSize.Scale(14.0 / 16.0),
		// },
		TxColor:       t.Colors["Dark"],
		shaper:        t.Shaper,
		Command:       p.Command,
		Header:        p.Header,
		HeaderBgColor: t.Colors["Primary"],
		HeaderPadding: p.HeaderPadding,
		Border:        p.Border,
		BorderColor:   p.BorderColor,
		Body:          p.Body,
		BodyBgColor:   p.BodyBgColor,
		BodyPadding:   p.BodyPadding,
		Footer:        p.Footer,
		FooterBgColor: t.Colors["Secondary"],
		FooterPadding: p.FooterPadding,
	}
}

func (p DuoUIpage) Layout(gtx layout.Context) {
	layout.Flex{
		Axis: layout.Vertical,
	}.Layout(gtx,
		layout.Rigid(pageElementLayout(gtx, layout.N, p.HeaderBgColor, p.HeaderPadding, p.Header)),
		layout.Flexed(1, func(gtx layout.Context) layout.Dimensions {
			//DuoUIfill(gtx, p.BorderColor)
			layout.UniformInset(unit.Dp(p.Border)).Layout(gtx, pageElementLayout(gtx, layout.N, p.BodyBgColor, p.BodyPadding, p.Body))
			return layout.Dimensions{}

		}),
		layout.Rigid(pageElementLayout(gtx, layout.N, p.FooterBgColor, p.FooterPadding, p.Footer)),
	)
}

func pageElementLayout(gtx layout.Context, direction layout.Direction, background string, padding float32, elementContent func(gtx layout.Context) layout.Dimensions) func(gtx layout.Context) layout.Dimensions {
	return func(gtx layout.Context) layout.Dimensions {
		//hmin := gtx.Constraints.Width.Max
		//vmin := gtx.Constraints.Height.Min
		layout.Stack{Alignment: layout.W}.Layout(gtx,
			layout.Expanded(func(gtx layout.Context) layout.Dimensions {
				//rr := float32(gtx.Px(unit.Dp(0)))
				//clip.Rect{
				//	Rect: f32.Rectangle{Max: f32.Point{
				//		//X: float32(gtx.Constraints.Width.Min),
				//		//Y: float32(gtx.Constraints.Height.Min),
				//	}},
				//	NE: rr, NW: rr, SE: rr, SW: rr,
				//}.Op(gtx.Ops).Add(gtx.Ops)
				//fill(gtx, HexARGB(background))
				//pointer.Rect(image.Rectangle{Max: gtx.Dimensions.Size}).Add(gtx.Ops)
				return layout.Dimensions{}
			}),
			layout.Stacked(func(gtx layout.Context) layout.Dimensions {
				//gtx.Constraints.Width.Min = hmin
				//gtx.Constraints.Height.Min = vmin
				direction.Layout(gtx, func(gtx layout.Context) layout.Dimensions {
					layout.Flex{}.Layout(gtx,
						layout.Flexed(1, func(gtx layout.Context) layout.Dimensions {
							//layout.Inset{
							//	Top:    unit.Dp(padding),
							//	Right:  unit.Dp(padding),
							//	Bottom: unit.Dp(padding),
							//	Left:   unit.Dp(padding),
							//}.Layout(gtx, elementContent)
							return layout.Dimensions{}
						}))
					return layout.Dimensions{}
				})
				return layout.Dimensions{}
			}),
		)
		return layout.Dimensions{}
	}
}
