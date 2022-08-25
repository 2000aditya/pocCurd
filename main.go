package main

import (
	"pocCurd/db"
	users "pocCurd/handlers"
	user "pocCurd/models"

	"github.com/labstack/echo"
)

func main() {
	e := echo.New()

	d := db.DBConnect()
	h := users.NewHandler(user.NewUserModel(d))

	e.POST("/employee", h.SetDetail)
	e.PUT("/employee", h.UpdateDetail)
	e.DELETE("/employee/:id", h.DeleteDetails)
	e.GET("/employee", h.GetDetail)

	e.Logger.Fatal(e.Start(":1324"))
}
