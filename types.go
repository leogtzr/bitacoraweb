package main

import (
	"net/http"
)

const (
	connHost = "localhost"
	connPort = "8080"

	userEnvVar     = "BITACORA_USER"
	passwordEnvVar = "BITACORA_PASSWORD"

	enterYourUserNamePassword = "Please enter your username and password"
)

// Route ...
type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

// Routes ...
type Routes []Route

var (
	bitacoraUser     string
	bitacoraPassword string

	connectionError error

	routes = Routes{}
)

// BitacoraEntries ...
type BitacoraEntries struct {
	// Entries ...
	Entries []struct {
		Date        string `json:"date"`
		Description string `json:"description"`
	} `json:"entries"`
	Name     string `json:"name"`
	Nickname string `json:"nickname"`
}

const (
	// APIPort ...
	APIPort = "8083"
)
