package model

type Todo struct {
	BaseModel
	Name string `json:"name"`
}
