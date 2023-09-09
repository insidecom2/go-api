package interfaces

type User struct {
	Name    string `json:"name"`
	Surname string `json:"surname"`
}

type ResponseUser struct {
	Name    string
	Surname string
}
