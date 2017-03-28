package main

import (
	"fmt"
	"io"
	"net/http"
	"gowork/common/application"
	"gowork/common/convertor"
)

func HandlerHello(w http.ResponseWriter, req *http.Request) {
	fmt.Println("received an request...")
	io.WriteString( w, "Welcome to gowork!\n")
}

func main () {
	config := map[string]interface{}{}
	// set an value into map
    value := "hello, gowork !"
    config["say"] = value
	// new application using the map, which contains an key "say" and value "hello, gowork !"
    app := application.NewApplication("sample", config)

	// now we can get the value of key "say"
    v := app.Get("say")
    str, _ := convertor.ToString(v) // you can get value of "say" is "hello, gowork !"
	fmt.Println(str)

	// set another key/value into the instance of application
	app.Set("another", 100)
	v = app.Get("another")
    vInt, _ := convertor.ToInt(v) // you can get value of "another" is 100
	fmt.Println(vInt)

	// add handler to serve as http server
	app.AddHandlerFunc("/hello", HandlerHello)

	// application run
    err := app.Go()
    if err != nil {
	    fmt.Println("Error: %s", err.Error())
    }
}