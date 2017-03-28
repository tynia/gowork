package main

import (
	"fmt"
	"io"
	"net/http"
	"gowork/common/application"
)

func HandlerHello(w http.ResponseWriter, req *http.Request) {
	fmt.Println("received an request...")
	io.WriteString( w, "Welcome to gowork!\n")
}

func main () {
	config := map[string]interface{}{}
	app := application.NewApplication("sample", config)

	app.AddHandlerFunc("/hello", HandlerHello)

	err := app.Go()
	if err != nil {
		fmt.Println("Error: %s", err.Error())
	}
}