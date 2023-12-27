package formclient

import (
	"syscall/js"
)

func (f *FormClient) validateForm(source_input *js.Value) (err string) {

	// 1 chequear input origen
	source_field_name := source_input.Get("name").String()

	source_fields, err := f.ObjectActual().GetFieldsByNames(source_field_name)
	if err != "" {
		return
	}

	input, new_value, err := f.getFormInputValue(&source_fields[0])
	if err != "" {
		return
	}

	f.setActionTypeFormData()

	// f.Log("f.its_new:", f.its_new)
	// f.Log("f.its_update_or_delete:", f.its_update_or_delete)

	f.err = f.fieldCheck(&source_fields[0], &input, new_value)
	if f.err != "" {
		f.notify(f.err)
		return
	}

	err = f.ObjectActual().ValidateData(f.its_new, f.its_update_or_delete, f.ObjectActual().FormData)

	f.Log("* RESUMEN FORMULARIO OK:", f.ObjectActual().FormData)

	f.notify(err)

	return
}

func (f *FormClient) notify(err string) {
	if f.ObjectActual().FrontHandler.FormNotify != nil {

		if err != "" {

			f.ObjectActual().FrontHandler.NotifyFormERR()

		} else {
			f.ObjectActual().FrontHandler.NotifyFormIsOK()

		}

	} else if err == "" && f.its_new {

		f.Log(f.ObjectActual().ObjectName, "no tiene FormNotify creamos la vista err:", err)
		// si no hay error y es de tipo nuevo

		err = f.CreateObjectsInDB(f.ObjectActual().Table, true, f.ObjectActual().FormData)
		if err != "" {
			f.UserMessage(err)
		} else {

			// r.Log("en db ", f.ObjectActual().FormData)
			err = f.ObjectActual().FrontHandler.SetObjectInDomAfterCreate(f.ObjectActual().FormData)
			if err != "" {
				f.UserMessage(err)
			} else { // click en el elemento creado

				f.Log(f.ObjectActual().ClickingID())

			}
		}

	}

}
