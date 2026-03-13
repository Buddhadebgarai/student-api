package types

type Student struct {
	ID   int    `json:"id"`
	Name string ` validate:"required" json:"name"`
	Age  int    `json:"age"`
}
