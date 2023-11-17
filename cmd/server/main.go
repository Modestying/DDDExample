package main

import "ddd_demo/application"

func main() {
	app := application.NewBankApp()
	app.Run()
}
