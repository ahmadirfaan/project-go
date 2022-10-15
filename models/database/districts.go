package database

type Districts struct {
	Id        *string `json:"id" gorm:"type:char(7)" json:"-"`
	RegencyId string  `json:"regencyId"`
	Name      string  `json:"name"`
}
