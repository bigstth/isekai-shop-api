package model

type (
	ItemCreatingReq struct {
		AdminID     string
		Name        string `json:"name" validate:"required,max=100"`
		Description string `json:"description" validate:"required,max=500"`
		Picture     string `json:"picture" validate:"required,max=500"`
		Price       uint   `json:"price" validate:"required"`
	}

	ItemEditingReq struct {
		AdminID     string
		Name        string `json:"name" validate:"required,max=100"`
		Description string `json:"description" validate:"required,max=500"`
		Picture     string `json:"picture" validate:"required,max=500"`
		Price       uint   `json:"price" validate:"required"`
	}
)
