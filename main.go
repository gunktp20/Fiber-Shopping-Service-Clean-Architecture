package main

import (
	"context"
	"shopping-service-be/pkg/config"
	"shopping-service-be/pkg/database"
	"shopping-service-be/server"
)

func main() {
	ctx := context.Background()
	_ = ctx

	conf := config.GetConfig()

	db := database.NewPostgresDatabase(ctx, conf)

	server.NewFiberServer(conf, db).Start()
}
