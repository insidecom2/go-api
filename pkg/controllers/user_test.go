package controllers

import (
	"testing"
)

func TestRequestCreateUserValidator(t *testing.T){
	var u *UserInterface
	// v := validator.New()
	// u.user.Name =""
	err := u.Validate()

	if err != nil {
		t.Fatal(err)
	}

}