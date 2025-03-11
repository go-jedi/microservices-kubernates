package posts

import (
	"github.com/go-jedi/gateway/internal/client"
	"github.com/go-jedi/gateway/pkg/logger"
)

type Service struct {
	logger *logger.Logger
	client *client.Client
}

func New(
	logger *logger.Logger,
	client *client.Client,
) *Service {
	return &Service{
		logger: logger,
		client: client,
	}
}
