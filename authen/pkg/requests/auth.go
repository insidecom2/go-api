package requests


type LoginBody struct {
	Email string ` validate:"email,required"`
	Password string ` validate:"required"`
}