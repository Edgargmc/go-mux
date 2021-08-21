package main

import (
	"os"
	"strconv"
)


const defaultPort = "8080"

func main() {
	app := App{}
	app.Initialize(
		os.Getenv("APP_DB_HOST"),
		getEnvAsInt("APP_DB_PORT", 5432),
		os.Getenv("APP_DB_USERNAME"),
		os.Getenv("APP_DB_PASSWORD"),
		os.Getenv("APP_DB_NAME"),
	)
	app.Run(os.Getenv(getServerPort()))
}

func getServerPort() string {
	serverPort := os.Getenv("PORT")
	if serverPort == "" {
		serverPort = defaultPort
	}
	return serverPort
}

func getEnv(key string, defaultVal string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return defaultVal
}

func getEnvAsInt(name string, defaultVal int) int {
	valueStr := getEnv(name, "")
	if value, err := strconv.Atoi(valueStr); err == nil {
		return value
	}
	return defaultVal
}