package posts

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/go-jedi/posts/internal/domain/posts"
)

func (r *Repository) All(ctx context.Context) ([]posts.Posts, error) {
	r.logger.Debug("[get all posts] execute repository")

	ctxTimeout, cancel := context.WithTimeout(ctx, time.Duration(r.db.QueryTimeout)*time.Second)
	defer cancel()

	q := `
		SELECT *
		FROM posts
		ORDER BY id
	`

	rows, err := r.db.Pool.Query(ctxTimeout, q)
	if err != nil {
		if errors.Is(err, context.DeadlineExceeded) {
			r.logger.Error("request timed out while get posts", "err", err)
			return nil, fmt.Errorf("the request timed out: %w", err)
		}
		r.logger.Error("failed to get posts", "err", err)
		return nil, fmt.Errorf("could not get posts: %w", err)
	}
	defer rows.Close()

	var p []posts.Posts

	for rows.Next() {
		var post posts.Posts

		if err := rows.Scan(
			&post.ID, &post.Title, &post.Description,
			&post.CreatedAt, &post.UpdatedAt,
		); err != nil {
			r.logger.Error("failed to scan row to get all posts", "err", err)
			return nil, fmt.Errorf("failed to scan row to get all posts: %w", err)
		}

		p = append(p, post)
	}

	if err := rows.Err(); err != nil {
		r.logger.Error("failed to get all posts", "err", rows.Err())
		return nil, fmt.Errorf("failed to get all posts: %w", err)
	}

	return p, nil
}
