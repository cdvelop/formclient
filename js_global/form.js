function getObjectIdFromForm(form) {
    return form.querySelector(`[name="` + form.dataset.field_id + `"]`)
}