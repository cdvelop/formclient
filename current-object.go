package formclient

import (
	"strconv"
	"syscall/js"
)

func (f *FormClient) currentObject(input []js.Value) (err string) {
	const this = "currentObject error "
	if len(input) != 1 {
		return this + "se esperaban 1 argumentos y se enviaron: " + strconv.Itoa(len(input))
	}

	f.form = input[0].Get("form")
	if !f.form.Truthy() {
		return this + "no se logro obtener formulario"
	}

	form_name := f.form.Get("name").String()

	return f.setNewFormObject(form_name)
}

func (f *FormClient) setNewFormObject(new_object_name string) (err string) {

	const t = "setNewFormObject error "

	if f.obj == nil || f.obj.ObjectName != new_object_name {

		// f.Log("formulario nuevo: " + new_object_name + ", anterior: " + f.obj.Name)

		f.obj, err = f.GetObjectByNameMainHandler(new_object_name)
		if err != "" {
			return t + err
		}

		err = f.setHtmlModule()
		if err != "" {
			return t + err
		}

		err = f.setHtmlForm()
		if err != "" {
			return t + err
		}

		// update object
	}

	// f.Log("*OBJETO ACTUAL FORMULARIO:", f.obj.Name)

	return ""
}

func (f *FormClient) setHtmlModule() (err string) {

	html, e := f.GetHtmlModule(f.obj.ModuleName)
	if e != "" {
		return e
	}
	var ok bool

	f.module, ok = html.(js.Value)
	if !ok {
		return "setHtmlModule error js.Value no fue enviado en GetHtmlModule"
	}

	return ""
}

func (f *FormClient) setHtmlForm() (err string) {

	f.form = f.module.Call("querySelector", `form[name="`+f.obj.ObjectName+`"]`)
	// form := module_html.Call("querySelector", "form", "#"+o.ObjectName)
	if !f.form.Truthy() {
		return "error no se logro obtener formulario " + f.obj.ObjectName
	}

	return ""
}
