package main

import (
	"embed"

	"github.com/developerRalsei/desktop-app-template/api"
	"github.com/developerRalsei/desktop-app-template/app"

	"github.com/wailsapp/wails/v2"
	"github.com/wailsapp/wails/v2/pkg/options"
	"github.com/wailsapp/wails/v2/pkg/options/assetserver"
)

//go:embed all:frontend/dist
var assets embed.FS

func main() {
	// Create a built in rest api
	Api := api.NewApi()
	Api.SetupRoutes()

	// Create an instance of the app structure
	App := app.NewApp()

	// Create application with options
	err := wails.Run(&options.App{
		Title:  "first-desktop-app",
		Width:  1024,
		Height: 768,
		AssetServer: &assetserver.Options{
			Assets: assets,
		},
		BackgroundColour: &options.RGBA{R: 27, G: 38, B: 54, A: 1},
		OnStartup:        App.Startup,
		Bind: []any{
			App,
		},
	})
	if err != nil {
		println("Error:", err.Error())
	}
}
