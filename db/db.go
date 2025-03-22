package db

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v5"
)

type Db struct {
	ctx    context.Context
	cancel context.CancelFunc
	cfg    *Config

	conn *pgx.Conn
}

type Config struct {
	ConnectString string `yaml:"connect_string"`
	Host          string `yaml:"host"`
	Port          uint16 `yaml:"port"`
	Dbname        string `yaml:"db_name"`
	SslMode       string `yaml:"ssl_mode"`
	User          string `yaml:"user"`
	Pass          string `yaml:"pass"`
	DriverName    string `yaml:"driver_name"`
}

func NewDbService(ctx context.Context, cfg *Config) (*Db, error) {
	fmt.Println("DB IS STARTING UP...")
	dbc := &Db{
		ctx: ctx,
		cfg: cfg,
	}
	var err error
	connStr := fmt.Sprintf(`user=%s password=%s host=%s dbname=%s sslmode=%s`,
		cfg.User, cfg.Pass, cfg.Host, cfg.Dbname, cfg.SslMode)
	//connStr := fmt.Sprintf("user=postgres password=penis host=localhost dbname=code_shelf sslmode=disable")
	dbc.conn, err = pgx.Connect(ctx, connStr)
	if err != nil {
		return &Db{}, err
	}
	fmt.Println("Connected successfully")
	return dbc, nil
}
