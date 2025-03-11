package posts

import (
	"github.com/go-jedi/posts/pkg/logger"
	"github.com/go-jedi/posts/pkg/postgres"
)

type Repository struct {
	logger *logger.Logger
	db     *postgres.Postgres
}

func New(logger *logger.Logger, db *postgres.Postgres) *Repository {
	return &Repository{
		logger: logger,
		db:     db,
	}
}
