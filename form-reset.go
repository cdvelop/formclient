package formclient

func (f *FormClient) FormClean(object_name ...string) (err string) {
	const e = "FormClean "

	for _, name := range object_name {

		err = f.SetActualObject(name)
		if err != "" {
			return e + err
		}

		err = f.setNewFormObject()
		if err != "" {
			return e + err
		}
	}

	return f.reset()
}

func (f *FormClient) reset() (err string) {

	f.form.Call("reset")

	// seteamos los valores del formulario
	err = f.ObjectActual().ResetInputsViewForm(f.form)
	if err != "" {
		return err
	}

	// f.Log("ESTADO FORMULARIO DESPUÉS DE RESET:", f.obj.FormData)

	f.resetActionType()

	return
}
