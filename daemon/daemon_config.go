package daemon

import (
	"fmt"

	"github.com/sasbury/mini"
)

func configureFromFile(fileName string) {
	cfg, err := mini.LoadConfiguration(fileName)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(cfg)
}
