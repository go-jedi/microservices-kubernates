package posts

import (
	"github.com/go-jedi/posts/internal/service/posts"
	"github.com/go-jedi/posts/pkg/logger"
	"github.com/go-jedi/posts/pkg/validator"
	"github.com/gofiber/fiber/v3"
)

type Handler struct {
	postsService *posts.Service
	logger       *logger.Logger
	validator    *validator.Validator
}

func New(
	postsService *posts.Service,
	app *fiber.App,
	logger *logger.Logger,
	validator *validator.Validator,
) *Handler {
	h := &Handler{
		postsService: postsService,
		logger:       logger,
		validator:    validator,
	}

	h.initRoutes(app)

	return h
}

func (h *Handler) initRoutes(app *fiber.App) {
	api := app.Group("/v1/posts")
	{
		api.Post("", h.create)
		api.Get("/all", h.all)
	}
}
