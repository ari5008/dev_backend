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
	const idxName = "idx_account_id_track_id"
	dropIdxSQL := fmt.Sprintf("DROP INDEX IF EXISTS %s", idxName)
	// Execute the SQL statement to drop the index
	if err := dbConn.Exec(dropIdxSQL).Error; err != nil {
		panic(err)
	}
	// Create SQL statement for adding the unique index
	createIdxSQL := fmt.Sprintf("CREATE UNIQUE INDEX %s ON likeflags (account_id, track_id)", idxName)
	// Execute the SQL statement to create the index
	if err := dbConn.Exec(createIdxSQL).Error; err != nil {
		panic(err)
	}

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
