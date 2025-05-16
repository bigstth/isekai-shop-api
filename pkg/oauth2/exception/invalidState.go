package exception

type InvalidStateException struct {
}

func (e *InvalidStateException) Error() string {
	return "invalid state"
}
