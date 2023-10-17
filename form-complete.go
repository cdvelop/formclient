package formclient

import (
	"strings"

	"github.com/cdvelop/model"
)

func (f FormClient) FormAutoFill(o *model.Object) error {

	test_data, err := o.TestData(1, true, false)
	if err != nil {
		return err
	}

	err = f.FormComplete(o, test_data[0])
	if err != nil {
		return err
	}

	return nil
}

func (f FormClient) FormComplete(o *model.Object, data map[string]string) error {

	if o == nil {
		return model.Error("formComplete object nil")
	}

	module_html, err := f.dom.GetHtmlModule(o.ModuleName)
	if err != nil {
		return err
	}

	form, err := f.getHtmlForm(*module_html, o)
	if err != nil {
		return err
	}

	err = f.reset(form, o)
	if err != nil {
		return err
	}

	for _, field := range o.RenderFields() {

		input, err := getFormInput(*form, field)
		if err != nil {
			return err
		}

		new_value := data[field.Name]

		// f.dom.Log("SELECCIÓN: ", field.Input.HtmlName(), field.Name, "VALOR:", new_value, input)

		switch field.Input.HtmlName() {
		case "checkbox":
			// Log("checkbox: ", f.Name, "tamaño", input.Length(), input)

			for i := 0; i < input.Length(); i++ {

				input_check := input.Index(i)

				value := input_check.Get("value").String()

				if strings.Contains(new_value, value) {
					input_check.Set("checked", true)
				}

				// Log("input check:", input_check, "value:", value)
			}

		case "radio":

			for i := 0; i < input.Length(); i++ {

				input_radio := input.Index(i)

				value := input_radio.Get("value").String()

				if value == new_value {
					input_radio.Set("checked", true)
					break
				}
			}

			// Log("SELECCIÓN radio: ", f.Name, input)

		default:

			input.Set("value", new_value)
		}

		// f.dom.Log("*** HASTA AQUÍ OK 1", field.Name)

	}

	return nil
}
