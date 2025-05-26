package exception

import "fmt"

type PlayerItemRemoving struct {
	ItemID uint64
}

func (e *PlayerItemRemoving) Error() string {
	return fmt.Sprintf("Failed to remove item with ID %d from player's inventory", e.ItemID)
}
