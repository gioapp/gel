package page

import (
	"gioui.org/layout"
	"gioui.org/text"
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
