package formclient

import (
	"syscall/js"
)

func (f *FormClient) userFormTyping(this js.Value, source_input []js.Value) interface{} {

	if f.timeout_typing.Truthy() {
		// Si hay un temporizador en curso, lo cancelamos
		js.Global().Call("clearTimeout", f.timeout_typing)
	}

	// Configuramos un nuevo temporizador para 500 milisegundos
	f.timeout_typing = js.Global().Call("setTimeout", js.FuncOf(func(this js.Value, null []js.Value) interface{} {

		// f.Log("ejecutando acción después de 500 milisegundos")

		err := f.currentObject(source_input)
		if err != nil {
			f.dom.Log(err)
			return nil
		}

		err = f.validateForm(&source_input[0])
		if err != nil {
			// f.Log(err.Error())
			return nil
		}

		f.setActionType()

		f.dom.Log("formulario correcto")

		// err = f.db.CreateObjectsInDB()

		return nil
	}), 500)

	return nil

}
