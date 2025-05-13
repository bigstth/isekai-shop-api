package exceptions

type AdminCreating struct {
	AdminID string
}

func (e *AdminCreating) Error() string {
	return "Admin with ID " + e.AdminID + " already exists"
}
