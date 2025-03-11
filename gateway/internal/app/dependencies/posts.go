package dependencies

import (
	postshandler "github.com/go-jedi/gateway/internal/adapter/http/handlers/posts"
	postsservice "github.com/go-jedi/gateway/internal/service/posts"
)

func (d *Dependencies) PostsService() *postsservice.Service {
	if d.postsService == nil {
		d.postsService = postsservice.New(
			d.logger,
			d.client,
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
