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
	dbConn.AutoMigrate(&model.User{}, &model.Account{}, &model.Track{}, &model.Likeflag{})

	useRepository := repository.NewUserRepository(dbConn)
	accountRepository := repository.NewAccountRepository(dbConn)
	trackRepository := repository.NewTrackRepository(dbConn)
	likeFlagRepository := repository.NewLikeFlagRepository(dbConn)

	userValidator := validator.NewUserValidator()
	accountValidator := validator.NewAccountValidator()

	userUsecase := usecase.NewUserUsecase(useRepository, userValidator)
	accountUsecase := usecase.NewAccountUsecase(accountRepository, accountValidator)
	trackUsecase := usecase.NewTrackUsecase(trackRepository)
	likeFlagUsecase := usecase.NewLikeFlagUsecase(likeFlagRepository)

	userController := controller.NewUserController(userUsecase, accountUsecase)
	accountController := controller.NewAccountController(accountUsecase)
	trackController := controller.NewTrackController(trackUsecase, accountUsecase)
	likeFlagController := controller.NewLikeFlagController(likeFlagUsecase, accountUsecase)

	e := router.NewRouter(userController, accountController, trackController, likeFlagController)
	e.Logger.Fatal(e.Start(":8080"))
}