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

func (f *FormClient) setFormData(o *model.Object, new_data map[string]string) {
	f.form_data = make(map[string]string, len(o.Fields))
	if new_data != nil {
		f.form_data = new_data
	}

	f.Log("***SET FORM DATA:", o.Name, new_data)

}

func (f *FormClient) FormComplete(o *model.Object, data map[string]string) error {

	if o == nil {
		return model.Error("FormComplete object nil")
	}

	err := f.SetNewFormObject(o.Name)
	if err != nil {
		return err
	}

	//reset data formulario
	f.setFormData(f.last_object, data)

	module_html, err := f.GetHtmlModule(o.ModuleName)
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
		case "file":
			if field.Input.ItemView != nil {

				object_id := data[o.PrimaryKeyName()]

				f.ReadStringDataAsyncInDB(model.ReadDBParams{
					FROM_TABLE:      "file",
					WHERE:           []string{"object_id"},
					SEARCH_ARGUMENT: object_id,
					ORDER_BY:        "",
					SORT_DESC:       false,
				}, func(new_data []map[string]string, err error) {

					if err != nil {
						f.Log(err)
						return
					}

					new_html := field.Input.BuildItemView(new_data...)
					// f.dom.Log("FILE INPUT HTML NUEVO:", new_html, "en input:", input)
					input.Set("innerHTML", new_html)
				})

			} else {
				f.Log(" ERROR ItemView nulo en FILE INPUT: ", o.Module.ModuleName, field.Name)
			}

		default:

			input.Set("value", new_value)
		}

		// f.Log("*** ", field.Name, " html name:", field.Input.HtmlName())
		// f.Log("*** value:", new_value)

	}

	return nil
}
