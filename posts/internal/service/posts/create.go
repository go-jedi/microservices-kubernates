package posts

import (
	"context"

	"github.com/go-jedi/posts/internal/domain/posts"
)

func (s *Service) Create(ctx context.Context, dto posts.CreateDTO) (posts.Posts, error) {
	s.logger.Debug("[create a new post] execute service")

	return s.postsRepository.Create(ctx, dto)
}
