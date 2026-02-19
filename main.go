package main

import (
	"embed"

	"cnb.cool/mliev/open/go-web/cmd"
	"github.com/muleiwu/gomander"
)

//go:embed templates/**
var templateFS embed.FS

//go:embed static/**
var staticFs embed.FS

func main() {

	thatFs := map[string]embed.FS{
		"templates":  templateFS,
		"web.static": staticFs,
	}

	gomander.Run(func() {
		cmd.Start(thatFs)
	})
}
