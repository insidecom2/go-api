package interfaces

type User struct {
	Name    string `json:"name" validate:"required"`
	Surname string `json:"surname" validate:"required"`
	Email string `json:"email" validate:"email,required"`
	Password string `json:"password" validate:"required"`
	Token string `json:"token"`
}


type ResponseUser struct {
	Name    string `json:"name"`
	Surname string `json:"surname"`
	Email string `json:"email"`
	Token string `json:"token"`
}
