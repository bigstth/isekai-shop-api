package exceptions

type PlayerCreating struct {
	PlayerID string
}

func (e *PlayerCreating) Error() string {
	return "Player with ID " + e.PlayerID + " already exists"
}
