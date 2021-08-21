package main

import "os"

func main() {
	app := App{}
	app.Initialize(
		os.Getenv("APP_DB_HOST"),
		5432,
		os.Getenv("APP_DB_USERNAME"),
		os.Getenv("APP_DB_PASSWORD"),
		os.Getenv("APP_DB_NAME"),
	)
	app.Run(":8080")
}
