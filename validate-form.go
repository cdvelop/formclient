package formclient

import (
	"syscall/js"
)

func (f *FormClient) validateForm(source_input *js.Value) (err string) {

	// 1 chequear input origen
	source_field_name := source_input.Get("name").String()

	source_fields, err := f.obj.GetFieldsByNames(source_field_name)
	if err != "" {
		return
	}

	input, new_value, err := f.getFormInputValue(&source_fields[0])
	if err != "" {
		return
	}

	err = f.fieldCheck(&source_fields[0], &input, new_value)
	if err != "" {
		return
	}

	f.setActionTypeFormData()

	err = f.obj.ValidateData(f.its_new, f.its_update_or_delete, f.obj.FormData)
	if err != "" {
		return
	}

	// // 2 chequear todos los inputs renderizados y solo del objeto origen
	// for _, field := range f.obj.FieldsToFormValidate() {

	// 	if field.Name != source_field_name {

	// 		input, new_value, e := f.getFormInputValue(&field)
	// 		if e != "" {
	// 			err = e
	// 			return
	// 		}

	// 		err = f.fieldCheck(&field, &input, new_value)
	// 		if err != "" {
	// 			return
	// 		}
	// 	}
	// }

	f.Log("* RESUMEN FORMULARIO OK:", f.obj.FormData)

	if f.obj.FrontHandler.NotifyFormComplete != nil {
		f.obj.FrontHandler.NotifyFormIsOK()
	}

	return
}
