package formclient

import (
	"syscall/js"

	"github.com/cdvelop/model"
)

func Add(h *model.MainHandler, o model.ObjectHandlerAdapter) {

	f := &FormClient{
		DataBaseAdapter:      h,
		MessageAdapter:       h.MessageAdapter,
		Logger:               h,
		ObjectHandlerAdapter: o,
		HtmlAdapter:          h,
		form:                 js.Value{},
		its_new:              false,
		its_update_or_delete: false,
		timeout_typing:       js.Value{},
	}
	h.FormAdapter = f

	js.Global().Set("userFormTyping", js.FuncOf(f.UserFormTyping))

}
