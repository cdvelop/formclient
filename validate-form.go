package formclient

import (
	"syscall/js"
)

func (f *FormClient) validateForm(source_input *js.Value) (err string) {
	const e = "validateForm error "
	// 1 chequear input origen
	source_field_name := source_input.Get("name").String()

	field, exist := f.object.FieldExist(source_field_name)
	if !exist {
		return e + "campo: " + source_field_name + " no existe en " + f.object.ObjectName
	}

	input, new_value, err := f.getFormInputValue(field)
	if err != "" {
		return
	}

	f.setActionTypeFormData()

	// f.Log("f.its_new:", f.its_new)
	// f.Log("f.its_update:", f.its_update)
	if field.IsPrimaryKey(f.object) && new_value == "" {
		return ""
	}

	err = inputRight(field, input, new_value)
	if err != "" {
		return e + err
	}

	// el campo se mantiene igual sin cambios
	if f.object.FormData[field.Name] == new_value {
		return ""
	}

	f.Log("---new value:", new_value, "campo:", field.Name)

	// campo cambio su valor actualizamos
	f.object.FormData[field.Name] = new_value

	// validamos todo el formulario

	return f.object.ValidateData(f.its_new, f.its_update, f.object.FormData)
}
