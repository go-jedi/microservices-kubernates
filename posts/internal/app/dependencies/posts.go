package dependencies

import (
	postshandler "github.com/go-jedi/posts/internal/adapter/http/handlers/posts"
	postsrepository "github.com/go-jedi/posts/internal/repository/posts"
	postsservice "github.com/go-jedi/posts/internal/service/posts"
)

func (d *Dependencies) PostsRepository() *postsrepository.Repository {
	if d.postsRepository == nil {
		d.postsRepository = postsrepository.New(
			d.logger,
			d.postgres,
		)
	}

	return d.postsRepository
}

func (d *Dependencies) PostsService() *postsservice.Service {
	if d.postsService == nil {
		d.postsService = postsservice.New(
			d.PostsRepository(),
			d.logger,
		)
	}

	return d.postsService
}

func (d *Dependencies) PostsHandler() *postshandler.Handler {
	if d.postsHandler == nil {
		d.postsHandler = postshandler.New(
			d.PostsService(),
			d.app,
			d.logger,
			d.validator,
		)
	}

	return d.postsHandler
}
