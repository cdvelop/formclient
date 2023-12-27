package formclient

import (
	"syscall/js"

	"github.com/cdvelop/strings"

	"github.com/cdvelop/model"
)

func (f *FormClient) FormComplete(object_name string, data map[string]string, validate, auto_grow bool) (err string) {
	const e = "FormComplete "

	f.err = f.SetActualObject(object_name)
	if f.err != "" {
		return e + f.err
	}

	if len(data) == 0 {
		return e + "no hay data para completar formulario en el objeto:" + f.ObjectActual().ObjectName
	}

	err = f.setNewFormObject()
	if err != "" {
		return e + err
	}

	// reseteamos solo los campos del formulario html
	f.reset()

	// f.Log("DATA PARA COMPLETAR FORMULARIO:", data)

	for _, field := range f.ObjectActual().RenderFields() {

		input, err := f.getFormInput(&field)
		if err != "" {
			return e + err
		}

		new_value := data[field.Name]
		// f.Log("SELECCIÓN OK: ", field.Input.HtmlName(), field.Name, "VALOR:", new_value, input)

		switch field.Input.HtmlName() {
		case "checkbox":
			// Log("checkbox: ", f.Name, "tamaño", input.Length(), input)

			for i := 0; i < input.Length(); i++ {

				input_check := input.Index(i)

				value := input_check.Get("value").String()

				if strings.Contains(new_value, value) != 0 {
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
		case "file":
			if field.Input.ItemViewAdapter != nil {
				object_id := data[f.ObjectActual().PrimaryKeyName()]

				if object_id != "" {

					f.ReadAsyncDataDB(model.ReadParams{
						FROM_TABLE: "file",
						WHERE:      []map[string]string{{"object_id": object_id}},
						ORDER_BY:   "",
						SORT_DESC:  false,
					}, func(r *model.ReadResults, err string) {

						if err != "" {
							f.Log(e + err)
							return
						}
						// f.Log("RESULTADO FILE ID object_id:", object_id, r.ResultsString)

						if len(r.ResultsString) != 0 {
							new_html := field.Input.BuildItemsView(r.ResultsString...)
							// f.Log("FILE INPUT HTML NUEVO:", new_html, "en input:", input)
							input.Set("innerHTML", new_html)
						}
					})
				}

			} else {
				return e + "nil ItemViewAdapter en FILE INPUT: " + f.ObjectActual().Module.ModuleName + " " + field.Name
			}
		case "textarea":
			input.Set("value", new_value)

			if auto_grow {
				_, err = f.ObjectActual().CallFunction("TextAreaAutoGrow", input)
				if err != "" {
					f.Log(e + err)
				}
			}

		default:
			input.Set("value", new_value)

		}

		// f.Log("*** ", field.Name, " html name:", field.Input.HtmlName(), "value:", new_value)

		if validate && new_value != "" {
			f.UserFormTyping(js.Global(), []js.Value{js.ValueOf(input)})
		}

	}

	return ""
}

// func (f FormClient) FormAutoFill(object_name string) (err string) {

// 	err = f.setNewFormObject(object_name)
// 	if err != "" {
// 		return
// 	}

// 	test_data, err := f.ObjectActual().TestData(1, true, false)
// 	if err != "" {
// 		return err
// 	}

// 	err = f.FormComplete(f.ObjectActual().ObjectName, test_data[0], false, false)
// 	if err != "" {
// 		return err
// 	}

// 	return ""
// }

// func (f *FormClient) setFormData(new_data map[string]string) {
// 	f.obj.FormData = make(map[string]string, 0)
// 	if new_data != nil {
// 		f.obj.FormData = new_data
// 	}
// 	// f.Log("***SET FORM DATA:", o.ObjectName, new_data)
// }
