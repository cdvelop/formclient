package formclient

import (
	"syscall/js"

	"github.com/cdvelop/model"
)

func Add(dom domAdapter, db model.DataBaseAdapter) *FormClient {

	f := FormClient{
		domAdapter:      dom,
		DataBaseAdapter: db,
		obj:             &model.Object{},
		html_form:       js.Value{},
		action_create:   false,
		action_update:   false,
		action_delete:   false,
		timeout_typing:  js.Value{},
	}

	return &f

}
