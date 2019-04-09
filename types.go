package main

import (
	"database/sql"
	"net/http"
)

const (
	connHost = "localhost"
	connPort = "8080"

	userEnvVar       = "CBE_USER"
	passwordEnvVar   = "CBE_PASSWORD"
	userDBEnvVar     = "DB_USER"
	passwordDBEnvVar = "DB_PASSWORD"
	dbNameEnvVar     = "DB_NAME"
	driverName       = "mysql"

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
	cbeUser     string
	cbePassword string
	// DB variables:
	dbUser     string
	dbPassword string
	dbName     string

	db              *sql.DB
	connectionError error

	routes = Routes{}
)
