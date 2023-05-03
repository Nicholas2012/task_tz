package main

import (
	"database/sql"
	ihttp "github.com/Nicholas2012/task_tz/internal/http"
	"github.com/Nicholas2012/task_tz/internal/service"
	"github.com/Nicholas2012/task_tz/internal/storage/postgres"
	"github.com/Nicholas2012/task_tz/pkg/randomuser"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"os"

	_ "github.com/lib/pq"
)

type Config struct {
	ListenAddr  string
	PostgresDSN string
}

func NewConfig() *Config {
	return &Config{
		ListenAddr:  getOrDefault("LISTEN_ADDR", ":8080"),
		PostgresDSN: getOrDefault("POSTGRES_DSN", ""),
	}
}

func main() {
	config := NewConfig()

	if config.PostgresDSN == "" {
		log.Fatal("POSTGRES_DSN is not set")
	}

	// init database
	db, err := sql.Open("postgres", config.PostgresDSN)
	if err != nil {
		log.Fatalf("open db: %v", err)
	}

	if db.Ping() != nil {
		log.Fatalf("ping db: %v", err)
	}

	// apply migrations
	if err := postgres.ApplyMigrations(db); err != nil {
		log.Fatalf("apply migrations: %v", err)
	}

	// init internal dependencies
	repo := postgres.New(db)
	randomuserCli := randomuser.New()
	svc := service.New(repo, randomuserCli)

	// init http handlers
	router := gin.Default()
	handlers := ihttp.New(svc)
	handlers.Register(router)

	// init http server
	log.Printf("starting server at %s", config.ListenAddr)
	if err := http.ListenAndServe(config.ListenAddr, router); err != nil {
		log.Fatalf("listen and serve: %v", err)
	}
}

func getOrDefault(key, def string) string {
	v := os.Getenv(key)
	if v == "" {
		return def
	}
	return v
}
