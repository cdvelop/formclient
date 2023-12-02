package formclient

import (
	"syscall/js"
)

func (f *FormClient) FormReset(object_name string) (err string) {
	const this = "FormReset "
	err = f.setNewFormObject(object_name)
	if err != "" {
		return this + err
	}

	module_html, err := f.getHtmlModule()
	if err != "" {
		return this + err
	}

	form, err := f.getHtmlForm(module_html)
	if err != "" {
		return this + err
	}

	return f.reset(form)
}

func (f *FormClient) reset(form *js.Value) (err string) {

	form.Call("reset")

	// seteamos los valores del formulario
	err = f.obj.ResetFormValues(true)
	if err != "" {
		return err
	}

	f.Log("ESTADO FORMULARIO DESPUÉS DE RESET:", f.obj.FormData)

	f.resetActionType()

	return
}
