package exception

import "fmt"

type InventoryFillingException struct {
	PlayerID string
	ItemID   uint64
}

func (e *InventoryFillingException) Error() string {
	return fmt.Sprintf("Failed to fill inventory for player id: %s with item: %d", e.PlayerID, e.ItemID)
}
