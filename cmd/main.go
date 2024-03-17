package main

import (
	"github.com/shafaalafghany/segokuning-social-app/config"
	"github.com/shafaalafghany/segokuning-social-app/internal/app"
)

func main() {
	cfg := config.NewConfig()
	app.Run(cfg)
}
