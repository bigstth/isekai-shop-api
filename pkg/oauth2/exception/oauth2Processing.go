package exception

type OAuth2ProcessingException struct {
}

func (e *OAuth2ProcessingException) Error() string {
	return "oauth2 processing failed"
}
