package posts

import (
	"context"

	"github.com/go-jedi/posts/internal/domain/posts"
)

func (s *Service) All(ctx context.Context) ([]posts.Posts, error) {
	s.logger.Debug("[get all posts] execute service")

	return s.postsRepository.All(ctx)
}
