package main

import "github.com/nikp359/go-exercises/kata/http-base-server/server"

func main() {
	app := server.NewApp("127.0.0.1:8080")
	app.Run()
}
