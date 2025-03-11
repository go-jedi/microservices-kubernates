package posts

import (
	"context"

	"github.com/go-jedi/gateway/internal/domain/posts"
)

func (s *Service) Create(ctx context.Context, dto posts.CreateDTO) (posts.Posts, error) {
	s.logger.Debug("[create a new post] execute service")

	return s.client.Posts.Create(ctx, dto)
}
