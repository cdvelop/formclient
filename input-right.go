package formclient

import (
	"syscall/js"

	"github.com/cdvelop/model"
)

func inputRight(field *model.Field, input js.Value, new_value string) error {

	data_option := input.Get("dataset").Get("option").String()

	err := field.Input.Validate.ValidateField(new_value, field.SkipCompletionAllowed, data_option)
	if err == nil {

		// f.Log("value: ", new_value, " input: ", input)

		if new_value != "" {
			js.Global().Call("inputRight", input)
		} else {
			js.Global().Call("inputNormal", input)
		}

		return nil
	}

	js.Global().Call("inputWrong", input, err.Error())

	return model.Error("campo", field.Legend, "no valido", err.Error())
}
