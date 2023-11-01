package formclient

import (
	"syscall/js"

	"github.com/cdvelop/model"
)

func Add(dom domAdapter, db model.DataBaseAdapter) *FormClient {

	f := FormClient{
		domAdapter:      dom,
		DataBaseAdapter: db,
		last_object:     &model.Object{},
		html_form:       js.Value{},
		data_object:     map[string]string{},
		action_create:   false,
		action_update:   false,
		action_delete:   false,
		timeout_typing:  js.Value{},
	}

	return &f

}
