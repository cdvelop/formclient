package formclient

import (
	"strconv"
	"syscall/js"
)

func (f *FormClient) formPrepareFromInput(input []js.Value) (err string) {
	const e = "formPrepareFromInput "
	if len(input) != 1 {
		return e + "se esperaban 1 argumentos y se enviaron: " + strconv.Itoa(len(input))
	}

	f.form = input[0].Get("form")
	if !f.form.Truthy() {
		return e + "no se logro obtener formulario"
	}

	f.object, f.err = f.GetObjectBY(f.form.Get("name").String(), "")
	if f.err != "" {
		return e + f.err
	}

	return
}

func (f *FormClient) setNewFormObject() (err string) {

	const e = "setNewFormObject "

	err = f.setHtmlModule()
	if err != "" {
		return e + err
	}

	err = f.setHtmlForm()
	if err != "" {
		return e + err
	}

	// f.Log("*OBJETO ACTUAL FORMULARIO:", f.obj.Name)

	return ""
}

func (f *FormClient) setHtmlModule() (err string) {
	const e = " func setHtmlModule"

	f.html_any, err = f.GetHtmlModule(f.object.ModuleName)
	if err != "" {
		return err + e
	}
	var ok bool

	f.module, ok = f.html_any.(js.Value)
	if !ok {
		return "js.Value no fue enviado en GetHtmlModule" + e
	}

	return ""
}

func (f *FormClient) setHtmlForm() (err string) {

	f.form = f.module.Call("querySelector", `form[name="`+f.object.ObjectName+`"]`)
	// form := module_html.Call("querySelector", "form", "#"+o.ObjectName)
	if !f.form.Truthy() {
		return "no se logro obtener formulario " + f.object.ObjectName + " func setHtmlForm"
	}

	return ""
}
