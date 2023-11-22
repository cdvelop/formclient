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

	obj *model.Object

	html_form js.Value

	action_create bool
	action_update bool
	action_delete bool

	timeout_typing js.Value
}
