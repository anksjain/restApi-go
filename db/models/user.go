package models

type User struct {
	Id      int     `json:"id"`
	FName   string  `json:"fname"`
	City    string  `json:"city"`
	Married bool    `json:"married"`
	Phone   int64   `json:"phone"`
	Height  float32 `json:"height"`
}

type UserQuery struct {
	Id     *int
	Ids    []*int
	IdsMap map[int]bool
}
