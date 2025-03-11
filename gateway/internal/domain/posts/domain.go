package posts

import "time"

// Posts represents a posts in the system.
type Posts struct {
	ID          int       `json:"id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

// CreateDTO represents the data required to create a new post.
type CreateDTO struct {
	Title       string `json:"title" validate:"omitempty,min=1"`
	Description string `json:"description" validate:"omitempty,min=1"`
}
