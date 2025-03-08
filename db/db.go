package db

import "context"

type Db struct {
	ctx    context.Context
	cancel context.CancelFunc
	cfg    *Config
}

type Config struct {
	ConnectString string `yaml:"connect_string"` // строка подключения к БД
	Host          string `yaml:"host"`           // host БД
	Port          string `yaml:"port"`           // порт листенера БД
	Dbname        string `yaml:"db_name"`        // имя БД
	SslMode       string `yaml:"ssl_mode"`       // режим SSL
	User          string `yaml:"user"`           // пользователь для подключения к БД
	Pass          string `yaml:"pass"`           // пароль пользователя
	DriverName    string `yaml:"driver_name"`    // имя драйвера "postgres" | "pgx" | "godror"
}

func NewDbService(ctx context.Context, cfg *Config) (*Db, error) {
	return &Db{
		ctx: ctx,
		cfg: cfg,
	}, nil
}
