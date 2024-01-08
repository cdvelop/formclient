package formclient

func (f *FormClient) FormClean(object_name ...string) (err string) {
	const e = "FormClean "

	for _, name := range object_name {

		f.object, f.err = f.GetObjectBY(name, "")
		if err != "" {
			return e + f.err
		}

		f.err = f.setNewFormObject()
		if f.err != "" {
			return e + f.err
		}
	}

	return f.reset()
}

func (f *FormClient) reset() (err string) {

	f.form.Call("reset")

	// seteamos los valores del formulario
	err = f.object.ResetInputsViewForm(f.form)
	if err != "" {
		return err
	}

	// f.Log("ESTADO FORMULARIO DESPUÉS DE RESET:", f.obj.FormData)

	f.resetActionType()

	return
}
