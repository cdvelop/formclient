package formclient

import (
	"syscall/js"

	"github.com/cdvelop/model"
)

func (f FormClient) getHtmlForm(module_html js.Value, o *model.Object) (*js.Value, error) {

	form := module_html.Call("querySelector", `form[name="`+o.Name+`"]`)
	// form := module_html.Call("querySelector", "form", "#"+o.Name)
	if !form.Truthy() {
		return nil, model.Error("error no se logro obtener formulario", o.Name)
	}

	return &form, nil
}
