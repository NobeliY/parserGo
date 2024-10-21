package config

import (
	"os"

	"github.com/joho/godotenv"

	"github.com/NobeliY/parserGo/utils/helpers"
)


type Cfg struct{
	Urls string	
}

func MustCfg() *Cfg {
	const op = "internal.config.cfg.MustCfg"
	if err := godotenv.Load(".env"); err != nil {
		panic(helpers.SprintF(op, err))
	}
	return &Cfg{
		Urls: os.Getenv("URLS"),
	}
}