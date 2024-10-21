package main

import (
	"github.com/geziyor/geziyor"
	"github.com/geziyor/geziyor/export"

	"github.com/NobeliY/parserGo/internal/config"
	"github.com/NobeliY/parserGo/internal/custom"
)

func main() {
	config := config.MustCfg()
	geziyor.NewGeziyor(&geziyor.Options{
		StartURLs: []string{config.Urls},
		ParseFunc: custom.ParseUrls,
		Exporters: []export.Exporter{&export.JSON{}},
	}).Start()
}