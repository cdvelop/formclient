package formclient

import (
	"syscall/js"
)

func (f *FormClient) getHtmlModule(module_name string) (v js.Value, err string) {

	html, e := f.GetHtmlModule(module_name)
	if e != "" {
		return js.Value{}, e
	}

	module_html, ok := html.(js.Value)
	if !ok {
		return js.Value{}, "FormComplete error js.Value no fue enviado en GetHtmlModule"
	}

	return module_html, ""
}
