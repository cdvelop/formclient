package formclient

import (
	"github.com/cdvelop/strings"

	"github.com/cdvelop/model"
)

func (f FormClient) FormAutoFill(object_name string) (err string) {

	err = f.setNewFormObject(object_name)
	if err != "" {
		return
	}

	test_data, err := f.obj.TestData(1, true, false)
	if err != "" {
		return err
	}

	err = f.FormComplete(f.obj.ObjectName, test_data[0])
	if err != "" {
		return err
	}

	return ""
}

func (f *FormClient) setFormData(new_data map[string]string) {
	f.obj.FormData = make(map[string]string, len(f.obj.Fields))
	if new_data != nil {
		f.obj.FormData = new_data
	}

	// f.Log("***SET FORM DATA:", o.ObjectName, new_data)

}

func (f *FormClient) FormComplete(object_name string, data map[string]string) (err string) {
	const this = "FormComplete error"
	if len(data) == 0 {
		return this + "no hay data enviada para completar formulario"
	}

	err = f.setNewFormObject(object_name)
	if err != "" {
		return this + err
	}

	//reset data formulario
	f.setFormData(data)

	// html, err := f.GetHtmlModule(o.ModuleName)
	// if err != nil {
	// 	return err
	// }

	// module_html,ok := html.(*js.Value)
	//  if !ok {
	// 	return "FormComplete error *js.Value no fue enviado en GetHtmlModule")
	//  }

	module_html, err := f.getHtmlModule()
	if err != "" {
		return this + err
	}

	form, err := f.getHtmlForm(module_html)
	if err != "" {
		return this + err
	}

	err = f.reset(form)
	if err != "" {
		return this + err
	}

	for _, field := range f.obj.RenderFields() {

		input, err := getFormInput(*form, field)
		if err != "" {
			return this + err
		}

		new_value := data[field.Name]

		// f.dom.Log("SELECCIÓN: ", field.Input.HtmlName(), field.Name, "VALOR:", new_value, input)

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
				object_id := data[f.obj.PrimaryKeyName()]
				if object_id != "" {

					f.ReadStringDataAsyncInDB(model.ReadDBParams{
						FROM_TABLE:      "file",
						WHERE:           []string{"object_id"},
						SEARCH_ARGUMENT: object_id,
						ORDER_BY:        "",
						SORT_DESC:       false,
					}, func(new_data []map[string]string, err string) {

						if err != "" {
							f.Log(this + err)
							return
						}

						new_html := field.Input.BuildItemsView(new_data...)
						// f.dom.Log("FILE INPUT HTML NUEVO:", new_html, "en input:", input)
						input.Set("innerHTML", new_html)
					})
				}

			} else {
				return this + "nil ItemViewAdapter en FILE INPUT: " + f.obj.Module.ModuleName + " " + field.Name
			}
		case "textarea":
			input.Set("value", new_value)

			err = f.obj.CallFunction("TextAreaAutoGrow", input)
			if err != "" {
				f.Log(this + err)
			}

		default:
			input.Set("value", new_value)

		}

		// f.Log("*** ", field.Name, " html name:", field.Input.HtmlName())
		// f.Log("*** value:", new_value)

	}

	return ""
}
