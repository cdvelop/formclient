package formclient

func (f *FormClient) setActionTypeFormData() {

	id, exist := f.object.FormData[f.object.PrimaryKeyName()]

	if exist && id != "" {

		f.Log("id existe y no este vació its_update id:", id)

		f.its_update = true

	} else {

		f.Log("no hay id its_new", id)

		f.its_new = true

	}
}

func (f *FormClient) resetActionType() {
	f.its_new = false
	f.its_update = false
}
