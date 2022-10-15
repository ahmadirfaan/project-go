package database

type Regencies struct {
	Id         *string `json:"id" redis:"id"`
	ProvinceId string  `json:"provinceId" redis:"provinceId"`
	Name       string  `json:"name" redis:"name"`
}
