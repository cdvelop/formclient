package formclient

import (
	"strconv"
	"syscall/js"
)

func (f *FormClient) prepareFormActual(input []js.Value) (err string) {
	const e = "prepareFormActual "
	if len(input) != 1 {
		return e + "se esperaban 1 argumentos y se enviaron: " + strconv.Itoa(len(input))
	}

	f.form = input[0].Get("form")
	if !f.form.Truthy() {
		return e + "no se logro obtener formulario"
	}

	f.err = f.SetActualObject(f.form.Get("name").String())
	if f.err != "" {
		return e + f.err
	}

	return f.setNewFormObject()
}

func (f *FormClient) setNewFormObject() (err string) {

	const e = "setNewFormObject "

	// if f.obj == nil || f.obj.ObjectName != new_object_name {

	// f.Log("formulario nuevo: " + new_object_name + ", anterior: " + f.obj.Name)

	// f.obj, err = f.ObjectActual(new_object_name)
	// if err != "" {
	// 	return t + err
	// }

	err = f.setHtmlModule()
	if err != "" {
		return e + err
	}

	err = f.setHtmlForm()
	if err != "" {
		return e + err
	}

	// update object
	// }

	// f.Log("*OBJETO ACTUAL FORMULARIO:", f.obj.Name)

	return ""
}

func (f *FormClient) setHtmlModule() (err string) {
	const e = " func setHtmlModule"

	f.html_any, err = f.GetHtmlModule(f.ObjectActual().ModuleName)
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

	f.form = f.module.Call("querySelector", `form[name="`+f.ObjectActual().ObjectName+`"]`)
	// form := module_html.Call("querySelector", "form", "#"+o.ObjectName)
	if !f.form.Truthy() {
		return "no se logro obtener formulario " + f.ObjectActual().ObjectName + " func setHtmlForm"
	}

	return ""
}
