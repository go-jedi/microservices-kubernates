package posts

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/go-jedi/posts/internal/domain/posts"
)

func (r *Repository) Create(ctx context.Context, dto posts.CreateDTO) (posts.Posts, error) {
	r.logger.Debug("[create a new post] execute repository")

	ctxTimeout, cancel := context.WithTimeout(ctx, time.Duration(r.db.QueryTimeout)*time.Second)
	defer cancel()

	q := `
		INSERT INTO posts(
		    title,
		    description
		) VALUES($1, $2)
		RETURNING *;
	`

	var np posts.Posts

	if err := r.db.Pool.QueryRow(
		ctxTimeout, q,
		dto.Title, dto.Description,
	).Scan(
		&np.ID, &np.Title, &np.Description,
		&np.CreatedAt, &np.UpdatedAt,
	); err != nil {
		if errors.Is(err, context.DeadlineExceeded) {
			r.logger.Error("request timed out while creating the post", "err", err)
			return posts.Posts{}, fmt.Errorf("the request timed out: %w", err)
		}
		r.logger.Error("failed to create post", "err", err)
		return posts.Posts{}, fmt.Errorf("could not create post: %w", err)
	}

	return np, nil
}
