package database

type Regencies struct {
	Id         *string `json:"id" json:"-"`
	ProvinceId string  `json:"provinceId"`
	Name       string  `json:"name"`
}
