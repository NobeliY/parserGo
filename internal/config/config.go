package config

import (
	"os"

	"github.com/NobeliY/parserGo/utils/helpers"
	"github.com/joho/godotenv"
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