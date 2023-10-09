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

	if f.last_object == nil {
		// log("primer inicio objeto id: " + form_name)
		return f.setCurrentObject(form_name)

	} else {

		if f.last_object.Name != form_name { //objeto ha cambiado
			f.dom.Log("objeto cambio nuevo: " + form_name + ", anterior: " + f.last_object.Name)

			//reset data formulario
			f.data_object = nil

			return f.setCurrentObject(form_name)
		}
	}

	return nil
}

func (f *FormClient) setCurrentObject(object_name string) error {

	object, err := f.dom.GetObjectByName(object_name)
	if err != nil {
		return err
	}

	f.last_object = object

	f.dom.Log("*OBJETO ACTUAL:", f.last_object.Name)

	return nil
}
