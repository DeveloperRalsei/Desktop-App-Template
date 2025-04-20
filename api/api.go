package api

import (
	"fmt"
	"net/http"

	"github.com/developerRalsei/desktop-app-template/api/handlers"
)

type Api struct{}

func NewApi() *Api {
	return &Api{}
}

func (a *Api) SetupRoutes() {
	http.HandleFunc("/api/hello", handlers.HelloHandler)
	go func() {
		http.ListenAndServe(":3000", nil)
		fmt.Println("Http server running on port 3000")
	}()
}
