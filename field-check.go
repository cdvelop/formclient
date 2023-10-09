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

	if f.data_object[field.Name] != new_value {
		f.dom.Log("---new value:", new_value, "campo:", field.Name)
		f.data_object[field.Name] = new_value
	}

	return nil
}
