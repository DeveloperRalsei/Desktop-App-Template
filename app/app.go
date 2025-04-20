package app

import (
	"context"
	"encoding/json"
	"os"

	"github.com/developerRalsei/desktop-app-template/config"
)

// define your app struct and bindings here

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
func (a *App) Startup(ctx context.Context) {
	a.ctx = ctx
}

func (a *App) GetAppConstants() (*config.AppConfig, error) {
	data, err := os.ReadFile("./wails.json")
	if err != nil {
		return nil, err
	}

	var config config.AppConfig
	if err = json.Unmarshal(data, &config); err != nil {
		return nil, err
	}

	return &config, nil
}

type db struct {
	Value int `json:"value"`
}

// takes "increase" or "decrease" or "" to get value
func (a *App) Counter(method string) (int, error) {
	file, err := os.OpenFile("./db.json", os.O_RDWR|os.O_CREATE, 0644)
	if err != nil {
		return 0, err
	}
	defer file.Close()

	var data db
	if err = json.NewDecoder(file).Decode(&data); err != nil && err.Error() != "EOF" {
		return 0, err
	}

	switch method {
	case "increase":
		data.Value += 1
	case "decrease":
		data.Value -= 1
	default:

	}

	file.Seek(0, 0)
	file.Truncate(0)

	err = json.NewEncoder(file).Encode(
		data,
	)
	if err != nil {
		return 0, err
	}
	return data.Value, nil
}
