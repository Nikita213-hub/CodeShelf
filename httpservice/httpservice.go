package httpservice

import (
	"context"
	"github.com/Nikita213-hub/CodeShelf/db"
	"github.com/Nikita213-hub/CodeShelf/middlewares"
	"net/http"
)

type Service struct {
	ctx      context.Context
	cancel   context.CancelFunc
	cfg      *Config
	Handlers Handlers
}

type Handler struct {
	Path       string
	HandleFunc http.HandlerFunc
	Name       string
}

type Handlers map[string]Handler

type Config struct {
	AuthType string `yaml:"auth_type"`
}

func New(ctx context.Context, cfg *Config, strg *db.Db) (*Service, error) {
	service := &Service{
		ctx: ctx,
		cfg: cfg,
	}
	service.Handlers = Handlers{
		"hello_handler": Handler{
			"/hello",
			middlewares.AuthMiddleware(helloHandler, strg),
			"hello_handler",
		},
		"signup": Handler{
			"/signup",
			SignUp(strg),
			"auth_handler",
		},
		"signin": Handler{
			"/signin",
			SignIn(strg),
			"signin",
		},
		"addSnippet": Handler{
			"/snippets/add",
			middlewares.AuthMiddleware(newSnippet(strg), strg),
			"add_snippet",
		},
		"getSnippet": Handler{
			"/snippets",
			getSnippet(strg),
			"get_snippet",
		},
		"uploadSnippet": Handler{
			"/snippets/upload",
			middlewares.AuthMiddleware(uploadSnippet(strg), strg),
			"upload_snippet",
		},
	}
	return service, nil
}
