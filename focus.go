package formclient

import (
	"github.com/cdvelop/model"
)

// focus_field_name default focus in first element
func (f FormClient) FormInputFocus(object_name string, focus_field_name ...string) (err string) {
	const e = "FormInputFocus "
	err = f.SetActualObject(object_name)
	if err != "" {
		return e + err
	}
	// f.Log(e, f.ObjectActual().ObjectName)

	var field_name string
	// buscamos si se envió algún nombre de campo especifico
	for _, name := range focus_field_name {
		field_name = name
	}

	// preparamos el formulario html
	err = f.setNewFormObject()
	if err != "" {
		return e + err
	}
	// f.Log("form:", f.form)

	var focus_field *model.Field

	for i, field := range f.ObjectActual().RenderFields() {

		// caso 1 si se envió campo especifico
		if field_name != "" && field.Name == field_name {
			focus_field = &field
			break
		}

		// caso 2 no se envió nada especifico tomamos el primer campo
		if field_name == "" && i == 0 {
			focus_field = &field
			break
		}
	}

	if focus_field == nil {
		return e + "no se encontró input pare realizar foco"
	}

	input, err := f.getFormInput(focus_field)
	if err != "" {
		return e + err
	}
	// f.Log("INPUT:", input)

	f.WaitFor(20, func() {
		result := input.Call("focus")
		if result.Truthy() {
			f.Log(e+" focus result:", result)
		}
	})

	return
}
