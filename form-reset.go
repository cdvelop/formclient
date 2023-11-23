package formclient

import (
	"syscall/js"

	"github.com/cdvelop/model"
)

func (f FormClient) FormReset(o *model.Object) (err string) {

	module_html, err := f.getHtmlModule(o.ModuleName)
	if err != "" {
		return err
	}

	form, err := f.getHtmlForm(module_html, o)
	if err != "" {
		return err
	}

	// seteamos los valores del formulario
	for k := range o.FormData {
		o.FormData[k] = ""
	}

	return f.reset(form, o)
}

func (f FormClient) reset(form *js.Value, o *model.Object) (err string) {

	form.Call("reset")

	for _, field := range o.RenderFields() {

		if field.Input.ResetViewAdapter != nil {
			err = field.Input.ResetAdapterView()
		}
	}

	f.Log("ESTADO FORMULARIO DESPUÉS DE RESET:", o.FormData)

	f.setActionType()

	return
}
