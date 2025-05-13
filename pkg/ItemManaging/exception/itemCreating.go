package exceptions

type ItemCreatingException struct {
}
type ItemAlreadyExistsException struct {
}

func (e *ItemCreatingException) Error() string {
	return "Item creating failed"
}
func (e *ItemAlreadyExistsException) Error() string {
	return "Item name already exists"
}
