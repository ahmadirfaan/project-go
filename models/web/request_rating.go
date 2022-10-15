package web

type RequestRating struct {
	Rating uint8 `json:"rating" validate:"required,min=1,max=5"`
}
