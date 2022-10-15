package database

type Provinces struct {
	Id   *string `json:"id" redis:"id"`
	Name string  `json:"name" redis:"name"`
}
