package internal

import (
	"github.com/gofiber/fiber/v2"
	"github.com/rs/xid"
)

func NewMovieController(ms *MovieService) *fiber.App {
	mc := &movieController{ms}
	app := fiber.New()
	app.Post("/create", mc.Create)
	app.Get("/search", mc.Query)
	app.Get("/:id", mc.Find)
	app.Delete("/:id", mc.Delete)
	return app
}

type movieController struct {
	MovieService *MovieService
}

func (mc *movieController) Create(c *fiber.Ctx) error {
	ms := &MovieSignature{}
	if err := c.BodyParser(ms); err != nil {
		return err
	}
	m, err := mc.MovieService.Create(ms.Movie, ms.Signature)
	if err != nil {
		return err
	}
	return c.JSON(m)
}

func (mc *movieController) Query(c *fiber.Ctx) error {
	term := c.Query("q")
	movies, err := mc.MovieService.Query(term)
	if err != nil {
		return err
	}
	return c.JSON(movies)
}

func (mc *movieController) Find(c *fiber.Ctx) error {
	id, err := xid.FromString(c.Params("id"))
	if err != nil {
		return fiber.NewError(400, "invalid id")
	}
	m, err := mc.MovieService.Find(id)
	if err != nil {
		return fiber.NewError(404, "not found")
	}
	return c.JSON(m)
}

func (mc *movieController) Delete(c *fiber.Ctx) error {
	id, err := xid.FromString(c.Params("id"))
	if err != nil {
		return fiber.NewError(400, "invalid id")
	}

	if err := mc.MovieService.Delete(id, c.Query("signature")); err != nil {
		return fiber.NewError(404, "not found")
	}
	return c.SendStatus(204)
}
