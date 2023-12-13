package formclient

func (f *FormClient) FormReset(object_name string) (err string) {
	const this = "FormReset "
	err = f.setNewFormObject(object_name)
	if err != "" {
		return this + err
	}

	return f.reset()
}

func (f *FormClient) reset() (err string) {

	f.form.Call("reset")

	// seteamos los valores del formulario
	err = f.obj.ResetFormValues(f.form, true)
	if err != "" {
		return err
	}

	// f.Log("ESTADO FORMULARIO DESPUÉS DE RESET:", f.obj.FormData)

	f.resetActionType()

	return
}
