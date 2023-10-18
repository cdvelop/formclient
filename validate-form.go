package formclient

import (
	"syscall/js"
)

func (f *FormClient) validateForm(source_input *js.Value) error {

	if f.data_object == nil {
		f.data_object = make(map[string]string, len(f.last_object.Fields))
	}

	// 1 chequear input origen
	source_field_name := source_input.Get("name").String()

	source_fields, err := f.last_object.GetFieldsByNames(source_field_name)
	if err != nil {
		return err
	}

	input, new_value, err := f.getFormInputValue(&source_fields[0])
	if err != nil {
		return err
	}

	err = f.fieldCheck(&source_fields[0], &input, new_value)
	if err != nil {
		return err
	}

	// 2 chequear todos los input menos origen
	for _, field := range f.last_object.Fields {

		if field.Name != source_field_name {

			input, new_value, err := f.getFormInputValue(&field)
			if err != nil {
				return err
			}

			err = f.fieldCheck(&field, &input, new_value)
			if err != nil {
				return err
			}
		}
	}

	f.dom.Log("*RESUMEN FORMULARIO:")
	for key, value := range f.data_object {
		f.dom.Log("FIELD NAME: ", key, " VALUE: ", value)
	}

	return nil
}
