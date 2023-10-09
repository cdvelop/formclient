package formclient

import (
	"strings"

	"github.com/cdvelop/model"
)

func (f FormClient) FormAutoFill(o *model.Object) {

	test_data, err := o.TestData(1, true, false)
	if err != nil {
		f.dom.Log(err)
	}

	err = f.formComplete(o, test_data[0])
	if err != nil {
		f.dom.Log(err)
	}

}

func (f FormClient) formComplete(o *model.Object, data map[string]string) error {

	if o == nil {
		return model.Error("formComplete object nil")
	}

	module_html, err := f.dom.GetHtmlModule(o.ModuleName)
	if err != nil {
		return err
	}

	form := module_html.Call("querySelector", `form[name="`+o.Name+`"]`)
	// form := module_html.Call("querySelector", "form", "#"+o.Name)
	if !form.Truthy() {
		return model.Error("formComplete error no se logro obtener formulario")
	}

	form.Call("reset")

	for _, f := range o.Fields {

		input, err := getHtmlInput(form, f)
		if err != nil {
			return err
		}

		new_value := data[f.Name]

		switch f.Input.HtmlName() {
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

			// Log("SELECCIÓN: ", f.Input.HtmlName(), f.Name, input)

			input.Set("value", new_value)
		}

	}

	return nil
}
