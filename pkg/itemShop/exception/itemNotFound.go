package exception

import "fmt"

type ItemNotFound struct {
	ItemID uint64
}

func (e *ItemNotFound) Error() string {
	return fmt.Sprintf("Item with ID %d not found", e.ItemID)
}
