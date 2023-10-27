package formclient

import (
	"syscall/js"

	"github.com/cdvelop/model"
)

func (d FormClient) getFormInputValue(field *model.Field) (input js.Value, value string, err error) {

	input = d.html_form.Get(field.Name)
	if !input.Truthy() {
		return js.Value{}, "", model.Error("input html", field.Name, "no encontrado")
	}
	var temp js.Value

	switch field.Input.HtmlName() {
	case "checkbox":
		var comma string
		// log("checkbox", field.Name)

		for i := 0; i < input.Length(); i++ {
			check := input.Index(i)
			temp = input.Index(i)

			if check.Get("checked").Bool() {

				value += comma + check.Get("value").String()
				comma = "," // se necesita coma para el siguiente elemento
			}
		}

		input = temp

	case "radio":
		// log("campo de tipo radio", field.Name)
		for i := 0; i < input.Length(); i++ {
			radio := input.Index(i)
			temp = input.Index(i)
			if radio.Get("checked").Bool() {
				value = radio.Get("value").String()
				break
			}
		}

		input = temp

	default:
		// log("campo de una sola entrada")
		value = input.Get("value").String()
	}

	return input, value, nil
}

func getFormInput(form js.Value, f model.Field) (js.Value, error) {

	if f.Input == nil {
		return js.Value{}, model.Error("getFormInput error. input nulo en campo", f.Name)
	}

	var input_type string
	var all string

	var input js.Value

	switch f.Input.HtmlName() {
	case "checkbox", "radio":
		input_type = "input[type='" + f.Input.HtmlName() + "']"
		all = "All"

	}

	input = form.Call("querySelector"+all, input_type+"[name='"+f.Name+"']")
	if !input.Truthy() {
		return js.Value{}, model.Error("input:", f.Name, "tipo:", f.Input.HtmlName(), "no encontrado en formulario")
	}

	return input, nil
}
