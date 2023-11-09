package formclient

import (
	"syscall/js"

	"github.com/cdvelop/model"
)

func (f *FormClient) fieldCheck(field *model.Field, input *js.Value, new_value string) error {

	if field.IsPrimaryKey(f.obj) && new_value == "" {
		return nil
	}

	err := inputRight(field, *input, new_value)
	if err != nil {
		return err
	}

	if f.obj.FormData[field.Name] != new_value {
		f.Log("---new value:", new_value, "campo:", field.Name)
		f.obj.FormData[field.Name] = new_value

		err := f.UpdateObjectsInDB(f.obj.Table, f.obj.FormData)
		if err != nil {
			return err
		}

		f.UserMessage("Registro Actualizado")

	}

	return nil
}
