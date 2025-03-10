package httpservice

import (
	"context"
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

func New(ctx context.Context, cfg *Config) (*Service, error) {
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
	}
	return service, nil
}
