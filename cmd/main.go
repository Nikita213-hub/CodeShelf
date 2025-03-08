package main

import (
	"context"
	"fmt"
	"github.com/Nikita213-hub/CodeShelf/config"
	"github.com/Nikita213-hub/CodeShelf/daemon"
)

func main() {
	cfg := config.NewCfg()
	err := cfg.LoadCfg()
	if err != nil {
		fmt.Println(err)
	}
	appDaemon, err := daemon.NewDaemon(context.Background(), cfg)
	if err != nil {
		fmt.Println(err)
	}
	appDaemon.Run()
}
