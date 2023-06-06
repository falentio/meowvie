package internal

import "github.com/gofiber/fiber/v2"

type searchController struct {
	SearchService *SearchService
}

func NewSearchController(ss *SearchService) *fiber.App {
	notNil(ss, "searchController.SearchService")
	sc := &searchController{ss}
	app := fiber.New()
	app.Get("/resync", sc.Resync)
	return app
}

func (s *searchController) Resync(c *fiber.Ctx) error {
	signature := c.Get("signature")
	return s.SearchService.Resync(signature)
}
