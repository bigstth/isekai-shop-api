package exceptions

type PlayerNotFound struct {
	PlayerID string
}

func (e *PlayerNotFound) Error() string {
	return "Player with ID " + e.PlayerID + " not found"
}
