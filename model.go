package formclient

import (
	"syscall/js"

	"github.com/cdvelop/model"
)

type FormClient struct {
	model.DataBaseAdapter
	model.MessageAdapter
	model.Logger
	model.ObjectsHandler
	model.HtmlAdapter

	obj *model.Object //objeto actual

	module js.Value //modulo html
	form   js.Value //formulario html

	its_new              bool
	its_update_or_delete bool

	timeout_typing js.Value
}
