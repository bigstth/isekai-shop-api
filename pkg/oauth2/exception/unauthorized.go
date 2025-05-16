package exception

type UnauthorizedException struct {
}

func (e *UnauthorizedException) Error() string {
	return "unauthorized"
}
