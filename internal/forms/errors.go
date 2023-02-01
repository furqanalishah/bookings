package forms

type errors map[string][]string

func (e errors) Add(field, message string) {
	e[field] = append(e[field], message)
}

func (e errors) Get(field string) string {
	errStr := e[field]
	if len(errStr) == 0 {
		return ""
	}
	return errStr[0]
}
