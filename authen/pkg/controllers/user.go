package controllers

import (
	"demoecho/pkg/response"
	"demoecho/pkg/services"
	"demoecho/pkg/utils"
	"fmt"
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var (
	userService    services.UserService
	profileService services.ProfileService
	profileClient  services.ProfileClient
)

type UserController interface {
	GetUser(c echo.Context) (err error)
}

type controllers struct{}

func NewUserController(servicesUser services.UserService) UserController {
	userService = servicesUser
	creds := insecure.NewCredentials()
	cc, err := grpc.Dial("localhost:50051", grpc.WithTransportCredentials(creds))
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Start connect grpc localhost:50051")

	profileClient = services.NewProfileClient(cc)
	profileService = services.NewProfileService(profileClient)

	return &controllers{}
}

func (con *controllers) GetUser(c echo.Context) (err error) {

	id := fmt.Sprint(c.Get("id"))
	fmt.Print(utils.GetContext(c, "id"))
	result, err := userService.GetUserService(id)
	profile, errProfile := profileService.GetProfile(id)

	if errProfile != nil {
		fmt.Printf("Cannot get profile %v", errProfile)
		return c.JSON(http.StatusBadRequest, response.ResponseFail("Cannot get profile"))
	}

	if err != nil {
		return c.JSON(http.StatusBadRequest, response.ResponseFail("Not found user"))
	}
	resUser := &response.ResponseUser{
		ID:      result.ID,
		Name:    result.Name,
		Surname: result.Surname,
		Email:   result.Email,
		Profile: profile,
	}
	return c.JSON(http.StatusOK, response.ResponseSuccess("OK", resUser))
}
