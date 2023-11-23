package formclient

import (
	"github.com/cdvelop/strings"

	"github.com/cdvelop/model"
)

func (f FormClient) FormAutoFill(o *model.Object) (err string) {

	test_data, err := o.TestData(1, true, false)
	if err != "" {
		return err
	}

	err = f.FormComplete(o, test_data[0])
	if err != "" {
		return err
	}

	return ""
}

func (f *FormClient) setFormData(o *model.Object, new_data map[string]string) {
	f.obj.FormData = make(map[string]string, len(o.Fields))
	if new_data != nil {
		f.obj.FormData = new_data
	}

	// f.Log("***SET FORM DATA:", o.ObjectName, new_data)

}

func (f *FormClient) FormComplete(o *model.Object, data map[string]string) (err string) {

	if o == nil {
		return "FormComplete object nil"
	}

	err = f.SetNewFormObject(o.ObjectName)
	if err != "" {
		return
	}

	//reset data formulario
	f.setFormData(f.obj, data)

	// html, err := f.GetHtmlModule(o.ModuleName)
	// if err != nil {
	// 	return err
	// }

	// module_html,ok := html.(*js.Value)
	//  if !ok {
	// 	return "FormComplete error *js.Value no fue enviado en GetHtmlModule")
	//  }

	module_html, err := f.getHtmlModule(o.ModuleName)
	if err != "" {
		return err
	}

	form, err := f.getHtmlForm(module_html, o)
	if err != "" {
		return err
	}

	err = f.reset(form, o)
	if err != "" {
		return err
	}

	for _, field := range o.RenderFields() {

		input, err := getFormInput(*form, field)
		if err != "" {
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

				object_id := data[o.PrimaryKeyName()]

				f.ReadStringDataAsyncInDB(model.ReadDBParams{
					FROM_TABLE:      "file",
					WHERE:           []string{"object_id"},
					SEARCH_ARGUMENT: object_id,
					ORDER_BY:        "",
					SORT_DESC:       false,
				}, func(new_data []map[string]string, err string) {

					if err != "" {
						f.Log(err)
						return
					}

					new_html := field.Input.BuildItemsView(new_data...)
					// f.dom.Log("FILE INPUT HTML NUEVO:", new_html, "en input:", input)
					input.Set("innerHTML", new_html)
				})

			} else {
				f.Log(" ERROR ItemViewAdapter nulo en FILE INPUT: ", o.Module.ModuleName, field.Name)
			}

		default:

			input.Set("value", new_value)
		}

		// f.Log("*** ", field.Name, " html name:", field.Input.HtmlName())
		// f.Log("*** value:", new_value)

	}

	return ""
}
