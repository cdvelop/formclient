package formclient

import (
	"syscall/js"

	"github.com/cdvelop/model"
)

func (f *FormClient) fieldCheck(field *model.Field, input *js.Value, new_value string) error {

	if field.IsPrimaryKey(f.last_object) && new_value == "" {
		return nil
	}

	err := inputRight(field, *input, new_value)
	if err != nil {
		return err
	}

	if f.form_data[field.Name] != new_value {
		f.Log("---new value:", new_value, "campo:", field.Name)
		f.form_data[field.Name] = new_value
		err := f.UpdateObjectsInDB(f.last_object.Table, f.form_data)
		if err != nil {
			return err
		}

		f.UserMessage("Registro Actualizado")

	}

	return nil
}
