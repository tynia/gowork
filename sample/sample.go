package main

import (
	"fmt"
	"github.com/tynia/gowork/service"
	"github.com/tynia/gowork/convertor"
	"io"
	"net/http"
)

func HandlerHello(w http.ResponseWriter, req *http.Request) {
	fmt.Println("received an request...")
	io.WriteString(w, "Welcome to gowork!\n")
}

func main() {
	config := map[string]string{}
	// set an value into map
	value := "hello, gowork !"
	config["say"] = value
	// new application using the map, which contains an key "say" and value "hello, gowork !"
	app := service.NewApplication("sample", config)

	// now we can get the value of key "say"
	v := app.Get("say")
	fmt.Println(v) // the console will print the value of "say" is "hello, gowork !"

	// set another key/value into the instance of application
	app.Set("another", "100")
	v = app.Get("another")
	vInt, _ := convertor.ToInt(v)
	fmt.Println(vInt) // the console will print the

	// add handler to serve as http server
	app.AddHandlerFunc("/hello", HandlerHello)

	// application run
	err := app.Go()
	if err != nil {
		fmt.Println("Error: %s", err.Error())
	}
}
