package database

import (
	"context"
	"fmt"
	"log"
	"shopping-service-be/pkg/config"
	"shopping-service-be/pkg/constant"
	"sync"
	"time"

	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/jmoiron/sqlx"
)

type postgresDatabase struct {
	Db *sqlx.DB
}

var (
	once       sync.Once
	dbInstance *postgresDatabase
)

func NewPostgresDatabase(pctx context.Context, conf *config.Config) *postgresDatabase {
	ctx, cancel := context.WithTimeout(pctx, 20*time.Second)
	defer cancel()

	once.Do(func() {
		dsn := fmt.Sprintf(
			"host=%s user=%s password=%s dbname=%s port=%d sslmode=%s TimeZone=%s",
			conf.Db.Host,
			conf.Db.User,
			conf.Db.Password,
			conf.Db.DBName,
			conf.Db.Port,
			conf.Db.SSLMode,
			conf.Db.TimeZone,
		)

		db, err := sqlx.Open("pgx", dsn)
		if err != nil {
			log.Fatalln(constant.Red+"Postgres database connection failed"+constant.Reset, err)
		}

		if err := db.PingContext(ctx); err != nil {
			log.Fatalln(constant.Red+"Postgres database connection failed"+constant.Reset, err)
		}

		dbInstance = &postgresDatabase{Db: db}
	})

	fmt.Println(constant.Green + "Postgres database connection successful" + constant.Reset)
	return dbInstance
}

func (p *postgresDatabase) GetDb() *sqlx.DB {
	return dbInstance.Db
}
