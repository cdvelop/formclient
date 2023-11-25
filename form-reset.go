package formclient

import (
	"syscall/js"
)

func (f *FormClient) FormReset(object_name string) (err string) {

	err = f.setNewFormObject(object_name)
	if err != "" {
		return
	}

	module_html, err := f.getHtmlModule()
	if err != "" {
		return err
	}

	form, err := f.getHtmlForm(module_html)
	if err != "" {
		return err
	}

	// seteamos los valores del formulario
	for k := range f.obj.FormData {
		f.obj.FormData[k] = ""
	}

	return f.reset(form)
}

func (f *FormClient) reset(form *js.Value) (err string) {

	form.Call("reset")

	for _, field := range f.obj.RenderFields() {

		if field.Input.ResetViewAdapter != nil {
			err = field.Input.ResetAdapterView()
		}
	}

	f.Log("ESTADO FORMULARIO DESPUÉS DE RESET:", f.obj.FormData)

	f.setActionType()

	return
}
