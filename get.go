package formclient

import (
	"syscall/js"

	"github.com/cdvelop/model"
)

func (f FormClient) getHtmlForm(module_html js.Value, o *model.Object) (jv *js.Value, err string) {

	form := module_html.Call("querySelector", `form[name="`+o.ObjectName+`"]`)
	// form := module_html.Call("querySelector", "form", "#"+o.ObjectName)
	if !form.Truthy() {
		return nil, "error no se logro obtener formulario " + o.ObjectName
	}

	return &form, ""
}
