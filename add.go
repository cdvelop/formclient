package formclient

import (
	"syscall/js"

	"github.com/cdvelop/model"
)

func Add(h *model.Handlers) (err string) {

	f := &FormClient{
		DataBaseAdapter:      h,
		MessageAdapter:       h.MessageAdapter,
		Logger:               h,
		ObjectsHandler:       h,
		HtmlAdapter:          h,
		obj:                  &model.Object{},
		html_form:            js.Value{},
		its_new:              false,
		its_update_or_delete: false,
		timeout_typing:       js.Value{},
	}
	h.FormAdapter = f

	err = h.CheckInterfaces("formclient config", *f)
	if err != "" {
		return err
	}

	js.Global().Set("userFormTyping", js.FuncOf(f.UserFormTyping))

	return ""

}
