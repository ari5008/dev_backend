package main

import (
	"backend/controller"
	"backend/db"
	"backend/model"
	"backend/repository"
	"backend/router"
	"backend/usecase"
	"fmt"
)

func main() {
	dbConn := db.NewDB()
	defer fmt.Println("successful")
	defer db.CloseDB(dbConn)
	dbConn.AutoMigrate(&model.User{})

	useRepository := repository.NewUserRepository(dbConn)
	userUsecase := usecase.NewUserUsecase(useRepository)
	userController := controller.NewUserController(userUsecase)
	e := router.NewRouter(userController)
	e.Logger.Fatal(e.Start(":8080"))
}