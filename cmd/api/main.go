package main

import (
	"log"
	"time"

	"github.com/BMPaiba/Go-Backend-Engineering/internal/db"
	"github.com/BMPaiba/Go-Backend-Engineering/internal/env"
	"github.com/BMPaiba/Go-Backend-Engineering/internal/store"
	"go.uber.org/zap"
)

//	@title			GopherSocial API
//	@description	This is a course.
//	@termsOfService	http://swagger.io/terms/

//	@contact.name	API Support
//	@contact.url	http://www.swagger.io/support
//	@contact.email	support@swagger.io

//	@license.name	Apache 2.0
//	@license.url	http://www.apache.org/licenses/LICENSE-2.0.html

// @BasePath					/v1
// @securityDefinitions.apikey	ApiKeyAuth
// @in							header
// @name						Authorization
// @description
func main() {
	log.SetFlags(log.Lshortfile)

	cfg := config{
		addr:   env.GetString("ADDR", ":3000"),
		apiURL: env.GetString("EXTERNAL_URL", "localhost:3000"),
		db: dbConfig{
			addr:         env.GetString("DB_ADDR", "postgres://admin:adminpassword@localhost/social?sslmode=disable"),
			maxOpenConns: env.GetInt("DB_MAX_OPEN_CONNS", 30),
			maxIdleConns: env.GetInt("DB_MAX_IDLE_CONNS", 30),
			maxIdleTime:  env.GetString("DB_MAX_IDLE_TIME", "15m"),
		},
		env: env.GetString("ENV", "development"),
		mail: mailConfig{
			exp: time.Hour * 24 * 3,
		},
	}

	// Logger
	logger := zap.Must(zap.NewProduction()).Sugar()
	defer logger.Sync()

	// Database
	database, err := db.New(cfg.db.addr, cfg.db.maxOpenConns, cfg.db.maxIdleConns, cfg.db.maxIdleTime)

	if err != nil {
		logger.Fatal(err)
	}

	defer database.Close()
	logger.Info("database conectada")

	store := store.NewStorage(database)

	app := &application{
		config: cfg,
		store:  store,
		logger: logger,
	}

	mux := app.mount()
	logger.Fatal(app.run(mux))
}
