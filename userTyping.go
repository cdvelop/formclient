package formclient

import (
	"syscall/js"
)

func (f *FormClient) UserFormTyping(this js.Value, source_input []js.Value) interface{} {

	if f.timeout_typing.Truthy() {
		// Si hay un temporizador en curso, lo cancelamos
		js.Global().Call("clearTimeout", f.timeout_typing)
	}

	// Configuramos un nuevo temporizador para 500 milisegundos
	f.timeout_typing = js.Global().Call("setTimeout", js.FuncOf(func(this js.Value, null []js.Value) interface{} {

		// f.Log("ejecutando acción después de 500 milisegundos")

		f.err = f.formPrepareFromInput(source_input)
		if f.err != "" {
			return f.Log(f.err)
		}

		f.err = f.validateForm(&source_input[0])
		if f.err != "" {
			return f.Log(f.err)
		}

		f.Log("formulario correcto")

		// err = f.db.CreateObjectsInDB()

		return nil
	}), 500)

	return nil

}

func (f *FormClient) SimulateUserFormTyping() {

	f.Log("SimulateUserFormTyping:", f.ObjectActual().ObjectName)

}
