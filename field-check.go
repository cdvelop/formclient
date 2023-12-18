package formclient

import (
	"syscall/js"

	"github.com/cdvelop/model"
)

func (f *FormClient) fieldCheck(field *model.Field, input *js.Value, new_value string) (err string) {

	if field.IsPrimaryKey(f.ObjectActual()) && new_value == "" {
		return ""
	}

	err = inputRight(field, *input, new_value)
	if err != "" {
		return err
	}

	if f.ObjectActual().FormData[field.Name] != new_value {
		f.Log("---new value:", new_value, "campo:", field.Name)
		f.ObjectActual().FormData[field.Name] = new_value

		if f.ObjectActual().FrontHandler.NotifyFormComplete == nil {

			err := f.UpdateObjectsInDB(f.ObjectActual().Table, f.ObjectActual().FormData)
			if err != "" {
				return err
			}

			f.UserMessage("Registro Actualizado")
		}

	}

	return ""
}
