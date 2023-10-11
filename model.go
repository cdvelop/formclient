package formclient

import (
	"syscall/js"

	"github.com/cdvelop/model"
)

type domAdapter interface {
	model.Log
	GetHtmlModule(module_name string) (*js.Value, error)
	GetObjectByName(name_to_search string) (*model.Object, error)
}

type FormClient struct {
	last_object *model.Object

	html_form   js.Value
	data_object map[string]string

	action_create bool
	action_update bool
	action_delete bool

	timeout_typing js.Value

	dom domAdapter
}
