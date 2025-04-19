package main

import (
	"context"
	"encoding/json"
	"os"
)

// App struct
type App struct {
	ctx context.Context
}

// NewApp creates a new App application struct
func NewApp() *App {
	return &App{}
}

// startup is called when the app starts. The context is saved
// so we can call the runtime methods
func (a *App) startup(ctx context.Context) {
	a.ctx = ctx
}

func (a *App) GetAppConstants() (*AppConfig, error) {
	data, err := os.ReadFile("./wails.json")
	if err != nil {
		return nil, err
	}

	var config AppConfig
	if err = json.Unmarshal(data, &config); err != nil {
		return nil, err
	}

	return &config, nil
}
