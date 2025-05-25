package exception

type PlayerCoinShowingException struct {
}

func (c *PlayerCoinShowingException) Error() string {
	return "Player coin showing failed."
}
