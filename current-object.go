package formclient

import (
	"strconv"
	"syscall/js"
)

func (f *FormClient) currentObject(input []js.Value) (err string) {
	const this = "currentObject error "
	if len(input) != 1 {
		return this + "se esperaban 1 argumentos y se enviaron: " + strconv.Itoa(len(input))
	}

	f.html_form = input[0].Get("form")
	if !f.html_form.Truthy() {
		return this + "no se logro obtener formulario"
	}

	form_name := f.html_form.Get("name").String()

	return f.setNewFormObject(form_name)
}

func (f *FormClient) setNewFormObject(new_object_name string) (err string) {

	if f.obj == nil || f.obj.ObjectName != new_object_name {

		// f.Log("formulario nuevo: " + new_object_name + ", anterior: " + f.obj.Name)

		object, err := f.GetObjectByName(new_object_name)
		if err != "" {
			return "setNewFormObject " + err
		}

		f.obj = object // update object
	}

	// f.Log("*OBJETO ACTUAL FORMULARIO:", f.obj.Name)

	return ""
}
