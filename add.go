package formclient

import (
	"syscall/js"

	"github.com/cdvelop/model"
)

func Add(h *model.MainHandler) {

	f := &FormClient{
		ObjectsHandlerAdapter: h,
		DataBaseAdapter:       h,
		MessageAdapter:        h,
		Logger:                h,
		HtmlAdapter:           h,
		DomAdapter:            h,
		ThemeAdapter:          h,
		TimeAdapter:           h,
		form:                  js.Value{},
		its_new:               false,
		its_update:            false,
		timeout_typing:        js.Value{},
	}
	h.FormAdapter = f

	js.Global().Set("userFormTyping", js.FuncOf(f.UserFormTyping))

}
