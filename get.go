package formclient

import (
	"syscall/js"
)

func (f FormClient) getHtmlForm(module_html js.Value) (jv *js.Value, err string) {

	form := module_html.Call("querySelector", `form[name="`+f.obj.ObjectName+`"]`)
	// form := module_html.Call("querySelector", "form", "#"+o.ObjectName)
	if !form.Truthy() {
		return nil, "error no se logro obtener formulario " + f.obj.ObjectName
	}

	return &form, ""
}
