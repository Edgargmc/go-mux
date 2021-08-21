package main

import "os"

func main() {
	app := App{}
	app.Initialize(
		os.Getenv("APP_DB_USERNAME_PROD"),
		os.Getenv("APP_DB_PASSWORD_PROD"),
		os.Getenv("APP_DB_NAME_PROD"),
	)
	app.Run(":8010")
}

//postgres
//edgargmc
//store
