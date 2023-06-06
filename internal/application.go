package internal

import (
	"fmt"
	"os"

	"github.com/blevesearch/bleve"
	"github.com/blevesearch/bleve/mapping"
	"github.com/caarlos0/env/v8"
	"github.com/goccy/go-json"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type Config struct {
	Port             string `env:"PORT" envDefault:"8080"`
	DatabaseUrl      string `env:"DATABASE_URL" envDefault:"file:./database/database.db?_journal=WAL"`
	SearchUrl        string `env:"SEARCH_URL" envDefault:"./database/bleve"`
	ApiSecret        string `env:"API_SECRET" envDefault:"secret"`
	AllowOrigins     string `env:"ALLOW_ORIGINS" envDefault:"*"`
	AutoMigrate      bool   `env:"AUTO_MIGRATE" envDefualt:"false"`
	LogflareSourceID string `env:"LOGFLARE_SOURCE_ID"`
	LogflareSecret   string `env:"LOGFLARE_SECRET"`
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
	f, _ := index.Fields()
	log.Debug().Any("f", f).Msg("field")
	db, err := gorm.Open(sqlite.Open(cfg.DatabaseUrl))
	if err != nil {
		panic("failed to open database, " + err.Error())
	}
	if cfg.AutoMigrate {
		err = db.Raw("delete from download_urls where coalesce(id, '') = ''").Error
		if err != nil {
			panic("failed to open database, " + err.Error())
		}
		if err := db.AutoMigrate(&Movie{}, &DownloadUrl{}); err != nil {
			panic("failed to do database migration, " + err.Error())
		}
	}
	if cfg.LogflareSecret != "" && cfg.LogflareSourceID != "" {
		logflare := NewZerologLogflare(cfg.LogflareSourceID, cfg.LogflareSecret)
		multi := zerolog.MultiLevelWriter(logflare, os.Stdout)
		log.Logger = log.Logger.Output(multi)
	}
	movieRepo := NewMovieRepoLru(NewMovieRepoGorm(db))
	downloadUrlRepo := NewDownloadUrlRepoGorm(db)

	ms := NewMovieService(movieRepo, downloadUrlRepo, search, signer)
	ss := NewSearchService(search, movieRepo, signer)

	app := fiber.New(fiber.Config{
		EnablePrintRoutes: true,
		JSONEncoder:       json.Marshal,
		JSONDecoder:       json.Unmarshal,
	})
	app.Use(recover.New())
	app.Use(RequestID)
	app.Use(Logger)
	app.Use(cors.New(cors.Config{
		MaxAge:       86400,
		AllowOrigins: cfg.AllowOrigins,
	}))
	app.Mount("movie", NewMovieController(ms))
	app.Mount("search", NewSearchController(ss))
	return &Application{app, cfg}
}
