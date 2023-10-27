package formclient

import (
	"syscall/js"

	"github.com/cdvelop/model"
)

func (f FormClient) FormReset(o *model.Object) error {

	module_html, err := f.dom.GetHtmlModule(o.ModuleName)
	if err != nil {
		return err
	}

	form, err := f.getHtmlForm(*module_html, o)
	if err != nil {
		return err
	}

	return f.reset(form, o)
}

func (f FormClient) reset(form *js.Value, o *model.Object) error {

	form.Call("reset")

	for _, field := range o.RenderFields() {

		if field.Input.InputReset != nil {
			field.Input.ResetInput()
		}
	}

	return nil
}
