package exception

type PlayerCoinShowingException struct {
}

func (c *PlayerCoinShowingException) Error() string {
	return "Player coins showing failed."
}
