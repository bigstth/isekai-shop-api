package exceptions

import "fmt"

type ItemArchivingException struct {
	ItemID uint64
}
type ItemNotFoundException struct {
	ItemID uint64
}

func (e *ItemArchivingException) Error() string {
	return fmt.Sprintf("Item archiving failed for item ID: %d", e.ItemID)
}

func (e *ItemNotFoundException) Error() string {
	return fmt.Sprintf("Item not found or already archived for item ID: %d", e.ItemID)
}
