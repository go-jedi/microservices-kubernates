package dependencies

import (
	postshandler "github.com/go-jedi/gateway/internal/adapter/http/handlers/posts"
	"github.com/go-jedi/gateway/internal/client"
	postsservice "github.com/go-jedi/gateway/internal/service/posts"
	"github.com/go-jedi/gateway/pkg/logger"
	"github.com/go-jedi/gateway/pkg/validator"
	"github.com/gofiber/fiber/v3"
)

type Dependencies struct {
	app       *fiber.App
	logger    *logger.Logger
	validator *validator.Validator
	client    *client.Client

	// posts
	postsService *postsservice.Service
	postsHandler *postshandler.Handler
}

func New(
	app *fiber.App,
	logger *logger.Logger,
	validator *validator.Validator,
	client *client.Client,
) *Dependencies {
	d := &Dependencies{
		app:       app,
		logger:    logger,
		validator: validator,
		client:    client,
	}

	d.initHandler()

	return d
}

// initHandler initialize handlers.
func (d *Dependencies) initHandler() {
	_ = d.PostsHandler()
}
