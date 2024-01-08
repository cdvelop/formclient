package formclient

import (
	"syscall/js"

	"github.com/cdvelop/model"
)

type FormClient struct {
	model.ObjectsHandlerAdapter
	model.DataBaseAdapter
	model.MessageAdapter
	model.Logger
	model.HtmlAdapter
	model.DomAdapter
	model.ThemeAdapter
	model.TimeAdapter

	// obj *model.Object //objeto actual
	html_any any

	module js.Value //modulo html
	form   js.Value //formulario html

	its_new    bool
	its_update bool

	object         *model.Object
	timeout_typing js.Value

	name_value string
	err        string
}
