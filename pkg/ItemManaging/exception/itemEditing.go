package exceptions

type ItemEditingException struct {
}

func (e *ItemEditingException) Error() string {
	return "Item editing failed"
}
