package formclient

import (
	"syscall/js"

	"github.com/cdvelop/model"
)

func (f *FormClient) getHtmlModule(module_name string) (js.Value, error) {

	html, err := f.GetHtmlModule(module_name)
	if err != nil {
		return js.Value{}, err
	}

	module_html, ok := html.(js.Value)
	if !ok {
		return js.Value{}, model.Error("FormComplete error js.Value no fue enviado en GetHtmlModule")
	}

	return module_html, nil
}
