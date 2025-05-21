package exception

type CoinAddingException struct {
}

func (c *CoinAddingException) Error() string {
	return "Coin adding failed"
}
