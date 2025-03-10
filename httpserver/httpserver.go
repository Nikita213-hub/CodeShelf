package httpserver

import (
	"context"
	"fmt"
	"github.com/Nikita213-hub/CodeShelf/httpservice"
	"net"
	"net/http"
	"time"
)

type HttpServer struct {
	ctx    context.Context
	cancel context.CancelFunc
	cfg    *Config

	listener net.Listener
	server   *http.Server
	Service  *httpservice.Service
	Mux      *http.ServeMux
}

type Config struct {
	ListenSpec      string        `yaml:"listen_spec" json:"listen_spec"`
	ReadTimeout     time.Duration `yaml:"read_timeout" json:"read_timeout"`
	WriteTimeout    time.Duration `yaml:"write_timeout" json:"write_timeout"`
	IdleTimeout     time.Duration `yaml:"idle_timeout" json:"idle_timeout"`
	ShutdownTimeout time.Duration `yaml:"shutdown_timeout" json:"shutdown_timeout"`
	MaxHeaderBytes  int           `yaml:"max_header_bytes" json:"max_header_bytes"`
	UseProfile      bool          `yaml:"use_go_profile" json:"use_profile"`
	UseTLS          bool          `yaml:"use_tls" json:"use_tls"`
	TLSCertFile     string        `yaml:"tls_cert_file" json:"tls_cert_file"`
	TLSKeyFile      string        `yaml:"tls_key_file" json:"tls_key_file"`
	TLSMinVersion   uint16        `yaml:"tls_min_version" json:"tls_min_version"`
	TLSMaxVersion   uint16        `yaml:"tls_max_version" json:"tls_max_version"`
}

func NewHttpServer(ctx context.Context, cfg *Config, service *httpservice.Service) (*HttpServer, error) {
	server := &http.Server{
		Addr: cfg.ListenSpec,
	}
	httpServer := &HttpServer{
		ctx:    ctx,
		cfg:    cfg,
		server: server,
	}
	httpServer.Mux = http.NewServeMux()
	httpServer.server.Handler = httpServer.Mux
	httpServer.Service = service
	for _, val := range httpServer.Service.Handlers {
		httpServer.Mux.HandleFunc(val.Path, val.HandleFunc)
		fmt.Println(val.Name, " is registered")
	}
	return httpServer, nil
}

func (s *HttpServer) Run() error {
	fmt.Println("Server is running...")
	return s.server.ListenAndServe()
}
