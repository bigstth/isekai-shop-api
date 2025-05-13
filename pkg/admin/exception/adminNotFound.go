package exceptions

type AdminNotFound struct {
	AdminID string
}

func (e *AdminNotFound) Error() string {
	return "Player with ID " + e.AdminID + " not found"
}
