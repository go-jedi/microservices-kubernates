package dependencies

import (
	postshandler "github.com/go-jedi/posts/internal/adapter/http/handlers/posts"
	postsrepository "github.com/go-jedi/posts/internal/repository/posts"
	postsservice "github.com/go-jedi/posts/internal/service/posts"
	"github.com/go-jedi/posts/pkg/logger"
	"github.com/go-jedi/posts/pkg/postgres"
	"github.com/go-jedi/posts/pkg/validator"
	"github.com/gofiber/fiber/v3"
)

type Dependencies struct {
	app       *fiber.App
	logger    *logger.Logger
	validator *validator.Validator
	postgres  *postgres.Postgres

	// posts
	postsRepository *postsrepository.Repository
	postsService    *postsservice.Service
	postsHandler    *postshandler.Handler
}

func New(
	app *fiber.App,
	logger *logger.Logger,
	validator *validator.Validator,
	postgres *postgres.Postgres,
) *Dependencies {
	d := &Dependencies{
		app:       app,
		logger:    logger,
		validator: validator,
		postgres:  postgres,
	}

	d.initHandler()

	return d
}

// initHandler initialize handlers.
func (d *Dependencies) initHandler() {
	_ = d.PostsHandler()
}
