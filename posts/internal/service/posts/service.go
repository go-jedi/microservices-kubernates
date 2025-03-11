package posts

import (
	"github.com/go-jedi/posts/internal/repository/posts"
	"github.com/go-jedi/posts/pkg/logger"
)

type Service struct {
	postsRepository *posts.Repository
	logger          *logger.Logger
}

func New(
	postsRepository *posts.Repository,
	logger *logger.Logger,
) *Service {
	return &Service{
		postsRepository: postsRepository,
		logger:          logger,
	}
}
