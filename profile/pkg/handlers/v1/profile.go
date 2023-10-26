package handlers

import "github.com/labstack/echo/v4"

type profileHandlers struct{}

type ProfileHandlers interface {
	GetProfile(c echo.Context) (err error)
	UpdateProfile(c echo.Context) (err error)
}

func NewProfileHandlers() ProfileHandlers {
	return &profileHandlers{}
}

func (p *profileHandlers) GetProfile(c echo.Context) (err error) {
	return nil
}

func (p *profileHandlers) UpdateProfile(c echo.Context) (err error) {
	return nil
}
