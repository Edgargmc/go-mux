package main

import "os"

func main() {
	a := App{}
	a.Initialize(
		os.Getenv("APP_DB_USERNAME_PROD"),
		os.Getenv("APP_DB_PASSWORD_PROD"),
		os.Getenv("APP_DB_NAME_PROD"),
		)

	a.Run(":8010")
}