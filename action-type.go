package formclient

func (f *FormClient) setActionType() {

	id, exist := f.form_data[f.last_object.PrimaryKeyName()]

	if exist {

		f.Log("id existe y no este vació:", id)
		f.action_create = false

		if !f.action_delete {
			f.action_update = true
			f.Log("acción es de tipo update")

		} else {
			f.Log("acción es de tipo delete")
		}

	} else {

		f.Log("no hay id es un objeto nuevo")

		f.action_create = true
		f.action_update = false
		f.action_delete = false
	}

}
