package exception

import "fmt"

type PlayerItemsFinding struct {
	PlayerID string
}

func (e *PlayerItemsFinding) Error() string {
	return fmt.Sprintf("Failed to find items for player %s", e.PlayerID)
}
