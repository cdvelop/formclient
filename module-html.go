package formclient

import (
	"syscall/js"
)

func (f *FormClient) getHtmlModule() (v js.Value, err string) {

	html, e := f.GetHtmlModule(f.obj.ModuleName)
	if e != "" {
		return js.Value{}, e
	}

	module_html, ok := html.(js.Value)
	if !ok {
		return js.Value{}, "getHtmlModule error js.Value no fue enviado en GetHtmlModule"
	}

	return module_html, ""
}
