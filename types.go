package main

type AppConfig struct {
	Name   string `json:"name"`
	Author struct {
		Name  string `json:"name"`
		Email string `json:"email"`
	} `json:"author"`
}
