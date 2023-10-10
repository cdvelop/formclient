package formclient

import (
	"syscall/js"

	"github.com/cdvelop/model"
)

func Add(d domAdapter) *FormClient {

	f := FormClient{
		last_object:    &model.Object{},
		html_form:      js.Value{},
		data_object:    map[string]string{},
		action_create:  false,
		action_update:  false,
		action_delete:  false,
		timeout_typing: js.Value{},
		dom:            d,
	}

	return &f

}
