package main

import (
	"context"

	"github.com/gunktp20/digital-hubx-be/pkg/config"
	"github.com/gunktp20/digital-hubx-be/pkg/database"
	"github.com/gunktp20/digital-hubx-be/server"
)

func main() {
	ctx := context.Background()
	_ = ctx

	conf := config.GetConfig()

	db := database.NewGormPostgresDatabase(ctx, conf)

	server.NewFiberServer(conf, db).Start()
}
