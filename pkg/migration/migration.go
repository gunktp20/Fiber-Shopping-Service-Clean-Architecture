package main

import (
	"context"
	"log"

	"github.com/gunktp20/digital-hubx-be/pkg/config"
	"github.com/gunktp20/digital-hubx-be/pkg/constant"
	"github.com/gunktp20/digital-hubx-be/pkg/database"
	"github.com/gunktp20/digital-hubx-be/pkg/models"
)

func main() {
	ctx := context.Background()
	conf := config.GetConfig()

	db := database.NewGormPostgresDatabase(ctx, conf).Db

	err := db.AutoMigrate(&models.AppGroup{}, &models.App{}, &models.Process{}, &models.SubProcess{}, &models.Type{})
	if err != nil {
		log.Fatalln(constant.Red+"Failed to migrate database : "+constant.Reset, err)
	} else {
		log.Fatalln(constant.Green + "Database migration successful" + constant.Reset)
	}

	defer func() {
		dbInstance, _ := db.DB()
		_ = dbInstance.Close()
	}()

}
