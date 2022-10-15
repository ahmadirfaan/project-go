package database

type Districts struct {
	Id        *string `json:"id" gorm:"type:char(7)" redis:"id"`
	RegencyId string  `json:"regencyId" redis:"regencyId"`
	Name      string  `json:"name"  redis:"name"`
}
