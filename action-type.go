package formclient

func (f *FormClient) setActionTypeFormData() {

	id, exist := f.ObjectActual().FormData[f.ObjectActual().PrimaryKeyName()]

	if exist && id != "" {

		f.Log("id existe y no este vació its_update_or_delete id:", id)

		f.its_update_or_delete = true

	} else {

		f.Log("no hay id its_new", id)

		f.its_new = true

	}
}

func (f *FormClient) resetActionType() {
	f.its_new = false
	f.its_update_or_delete = false
}
