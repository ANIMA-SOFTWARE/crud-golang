package types

type Orders struct {
}

type Order struct {
	ID    int    `json:"id" default:"0"`
	Name  string `json:"name" default:"test"`
	Email string `json:"email" default:"testemail"`
}
