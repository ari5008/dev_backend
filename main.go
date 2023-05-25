package main

import (
	"backend/controller"
	"backend/db"
	"backend/model"
	"backend/repository"
	"backend/router"
	"backend/usecase"
	"backend/validator"
	"fmt"
)

func main() {
	dbConn := db.NewDB()
	defer fmt.Println("successful")
	defer db.CloseDB(dbConn)
	dbConn.AutoMigrate(&model.User{}, &model.Account{})

	useRepository := repository.NewUserRepository(dbConn)
	accountRepository := repository.NewAccountRepository(dbConn)

	userValidator := validator.NewUserValidator()
	accountValidator := validator.NewAccountValidator()

	userUsecase := usecase.NewUserUsecase(useRepository, userValidator)
	accountUsecase := usecase.NewAccountUsecase(accountRepository, accountValidator)

	userController := controller.NewUserController(userUsecase, accountUsecase)
	accountController := controller.NewAccountController(accountUsecase)

	e := router.NewRouter(userController, accountController)
	e.Logger.Fatal(e.Start(":8080"))
}