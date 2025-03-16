package httpservice

import (
	"context"
	"github.com/Nikita213-hub/CodeShelf/db"
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
			helloHandler,
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
	}
	return service, nil
}
