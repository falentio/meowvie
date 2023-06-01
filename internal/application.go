package internal

import (
	"fmt"

	"github.com/blevesearch/bleve"
	"github.com/blevesearch/bleve/mapping"
	"github.com/caarlos0/env/v8"
	"github.com/goccy/go-json"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/gofiber/fiber/v2/middleware/requestid"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type Config struct {
	Port        string `env:"PORT" envDefault:"8080"`
	DatabaseUrl string `env:"DATABASE_URL" envDefault:"file:./database/database.db"`
	SearchUrl   string `env:"SEARCH_URL" envDefault:"./database/bleve"`
	ApiSecret   string `env:"API_SECRET" envDefault:"secret"`
}

type Application struct {
	App *fiber.App
	Cfg *Config
}

func (a *Application) Listen() error {
	addr := fmt.Sprintf(":%s", a.Cfg.Port)
	return a.App.Listen(addr)
}

func NewApplication() *Application {
	cfg := &Config{}
	if err := env.Parse(cfg); err != nil {
		panic("failed to get config from environment, " + err.Error())
	}

	signer := NewSigner(cfg.ApiSecret)

	index, err := bleve.New(cfg.SearchUrl, mapping.NewIndexMapping())
	if err != nil {
		index, err = bleve.Open(cfg.SearchUrl)
	}
	if err != nil {
		panic("failed to open bleve index, " + err.Error())
	}
	search := NewSearchBleve(index)

	db, err := gorm.Open(sqlite.Open(cfg.DatabaseUrl))
	if err != nil {
		panic("failed to open database, " + err.Error())
	}
	db.Raw("delete from download_urls where coalesce(id, '') = '';")
	if err := db.AutoMigrate(&Movie{}, &DownloadUrl{}); err != nil {
		panic("failed to do database migration, " + err.Error())
	}
	movieRepo := NewMovieRepoGorm(db)
	downloadUrlRepo := NewDownloadUrlRepoGorm(db)

	ms := NewMovieService(movieRepo, downloadUrlRepo, search, signer)

	app := fiber.New(fiber.Config{
		EnablePrintRoutes: true,
		JSONEncoder:       json.Marshal,
		JSONDecoder:       json.Unmarshal,
	})
	app.Use(recover.New())
	app.Use(requestid.New())
	app.Use(logger.New())
	app.Mount("movie", NewMovieController(ms))
	return &Application{app, cfg}
}
