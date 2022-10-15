package web

type ChangeTransactionRequest struct {
	StatusTransaction uint8 `json:"statusTransaction" validate:"required,min=1,max=3"`
}
