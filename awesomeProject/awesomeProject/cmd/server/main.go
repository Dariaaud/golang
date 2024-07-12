package main

import (
	"awesomeProject/accounts"
	"fmt"
	"os"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	accountsHandler := accounts.New()

	// Echo instance
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.GET("/account", accountsHandler.GetAccount)
	//e.POST("/account/create", accountsHandler.CreateAccount)

	// Start server
	command := os.Args[1]

	switch command {
	case "create":
		e.POST("/account/create", accountsHandler.CreateAccount)
	case "get":
		e.POST("/account/get", accountsHandler.GetAccount)
	case "delete":
		e.POST("/account/delete", accountsHandler.DeleteAccount)
	case "patch":
		e.POST("/account/create", accountsHandler.PatchAccount)
	case "change":
		e.POST("/account/create", accountsHandler.ChangeAccount)
	default:
		fmt.Println("Invalid command")
		os.Exit(1)
	}

	e.Logger.Fatal(e.Start(":1323"))
}
