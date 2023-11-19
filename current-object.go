package formclient

import (
	"strconv"
	"syscall/js"

	"github.com/cdvelop/model"
)

func (f *FormClient) currentObject(input []js.Value) error {

	if len(input) != 1 {
		return model.Error("en currentObject: se esperaban 1 argumentos y se enviaron:", strconv.Itoa(len(input)))
	}

	f.html_form = input[0].Get("form")
	if f.html_form.IsUndefined() {
		return model.Error("en currentObject: no se logro obtener formulario")
	}

	form_name := f.html_form.Get("name").String()

	return f.SetNewFormObject(form_name)
}

func (f *FormClient) SetNewFormObject(new_object_name string) error {

	if f.obj == nil || f.obj.ObjectName != new_object_name {

		// f.Log("formulario nuevo: " + new_object_name + ", anterior: " + f.obj.Name)

		object, err := f.GetObjectByName(new_object_name)
		if err != nil {
			return model.Error("SetNewFormObject", err)
		}

		f.obj = object // update object
	}

	// f.Log("*OBJETO ACTUAL FORMULARIO:", f.obj.Name)

	return nil
}
