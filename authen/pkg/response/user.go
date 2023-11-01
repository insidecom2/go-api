package response

import "demoecho/pkg/services"

type ResponseUser struct {
	ID      uint                         `json:"id"`
	Name    string                       `json:"name"`
	Surname string                       `json:"surname"`
	Email   string                       `json:"email"`
	Profile *services.GetProfileResponse `json:"profile"`
}
