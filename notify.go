package formclient

func (f *FormClient) notify(err string) {
	const e = "notify error "

	f.Log("* RESUMEN FORMULARIO OK:", f.object.FormData)

	if err != "" {
		f.Log("formulario correcto")
	}

	if f.object.FrontHandler.FormNotify != nil {

		if err != "" {
			f.object.FrontHandler.NotifyFormERR(err)
		} else {
			f.object.FrontHandler.NotifyFormIsOK()

		}

	} else if err == "" {

		f.Log(f.object.ObjectName, "no tiene FormNotify")

		if f.its_new {
			f.Log("** NUEVO OBJETO EN BASE DE DATOS LOCAL y REMOTO")

			err = f.CreateObjectsInDB(f.object.Table, true, f.object.FormData)
			if err != "" {
				f.UserMessage(e, "error", err)
				return
			}

			if f.object.FrontHandler.AfterCreate == nil {
				f.Log(e+"Objeto:", f.object.ObjectName, " no cuenta con FrontHandler.AfterCreate para actualizar vista")
				return
			}

			// ACTUALIZAMOS LA VISTA DEL OBJETO EN EL DOM
			err = f.object.FrontHandler.SetObjectInDomAfterCreate(f.object.FormData)
			if err != "" {
				f.UserMessage(e, err)
				return
			}

			// click en el elemento nuevo creado
			err = f.object.ClickingID()
			if err != "" {
				f.Log(e, err)
			}

			f.UserMessage("Nuevo Registro Ingresado")

		} else if f.its_update {
			f.Log("** ACTUALIZACIÓN EN BASE DE DATOS LOCAL y REMOTO")

			err := f.UpdateObjectsInDB(f.object.Table, true, f.object.FormData)
			if err != "" {
				f.UserMessage(e, err)
				return
			}

			if f.object.FrontHandler.AfterUpdate == nil {
				f.Log(e, "Objeto:", f.object.ObjectName, " no cuenta con FrontHandler.AfterUpdate para actualizar vista")
				return
			}

			// ACTUALIZAMOS LA VISTA DEL OBJETO EN EL DOM
			err = f.object.FrontHandler.SetObjectInDomAfterUpdate(f.object.FormData)
			if err != "" {
				f.UserMessage(e, err)
				return
			}

			f.UserMessage("Registro Actualizado")

		}

	}

}
