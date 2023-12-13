package main

import (
	"context"

	"github.com/sports-dynamics/cricket-fever/internal/app"
)

func main() {
	app.Start(context.Background(), app.Config{Port: 10000, ValidAPIToken: "sovApiToken"})
}
