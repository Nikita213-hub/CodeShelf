package daemon

import (
	"context"
	"errors"
	"fmt"

	"github.com/Nikita213-hub/CodeShelf/config"
	"github.com/Nikita213-hub/CodeShelf/db"
	"github.com/Nikita213-hub/CodeShelf/httpserver"
)

//Daemon initializes server, db, etc.

type Daemon struct {
	ctx    context.Context
	cancel context.CancelFunc
	cfg    *config.Config

	httpServer         *httpserver.HttpServer
	httpServerErrorsCh chan error
	dbService          *db.Db
	dbErrorsCh         chan error
}

func NewDaemon(ctx context.Context, cfg *config.Config) (*Daemon, error) {
	d := &Daemon{
		cfg:                cfg,
		httpServerErrorsCh: make(chan error),
		dbErrorsCh:         make(chan error),
	}
	if ctx == nil {
		d.ctx, d.cancel = context.WithCancel(context.Background())
	}
	d.ctx, d.cancel = context.WithCancel(ctx)
	server, err := httpserver.NewHttpServer(ctx, d.cfg.HttpServerCfg)
	if err != nil {
		return &Daemon{}, errors.New("error occurred while initializing daemon")
	}
	d.httpServer = server
	dbServ, err := db.NewDbService(ctx, d.cfg.DbCfg)
	if err != nil {
		return &Daemon{}, errors.New("error occurred while initializing daemon")
	}
	d.dbService = dbServ
	return d, nil
}

func (d *Daemon) Run() {
	go func() {
		d.httpServerErrorsCh <- d.httpServer.Run()
	}()

	for {
		var err error
		select {
		case err = <-d.httpServerErrorsCh:
			fmt.Println(err)
		}
	}
}
