# [gowork](https://github.com/tynia/gowork.git)
***GOWORK***, which sounds like **"狗窝"** in Chinese, is a tiny **server** framework implemented using golang.  
It is aimed at developing server faster.  
It is also an exercise to improve myself.  

# Author
**name**  : **tynia**  
**email** : [ekesi@live.com](ekesi@live.com)

# LICENSE
The **gowork** is released under the [**MIT LICENSE**](http://opensource.org/licenses/mit-license.php).  
you can import it without ANY concern on copyright.  

# Besides
I'd like you will leave my nickname(tynia) and [repository](https://github.com/tynia/gowork) in your project which you import the [**gowork**](https://github.com/tynia/gowork) in it.  

---
# Introduction
The **gowork** framework contains several modules, including:
- application 
- [logging](https://github.com/xuyu/logging)   --[**@xuyu**](https://github.com/xuyu) 

**[application]** 

It is a simple framework for server application development.  
When using the **gowork**, you need to do as follows:
```
import (
    "gowork/common/application"
    "gowork/common/application"
    ... // other imports you need
)

...

func main() {
    config := map[string]interface{}{}
    // set an value into map
    value := "hello, gowork !"
    config["say"] = value
    // new application using the map, which contains an key "say" and value "hello, gowork !"
    app := application.NewApplication("sample", config)

    // now we can get the value of key "say"
    v := app.Get("say")
    str, _ := convertor.ToString(v) 
    fmt.Println(str)   // the console will print the value of "say" is "hello, gowork !"

    // set another key/value into the instance of application
    app.Set("another", 100)
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
``` 
Above is the sample code to setup a server application, see the file: ```gowork/app/sample.go```  

---
In application.Go(), it will parse a json file. the file **must be** like this: 
```
{
	"Log": {
		"Level": "debug",
		"Suffix": "060102-15"
	},
	
	"Prog": {
		"CPU": 0,
		"Daemon": false,
		"HealthPort": "localhost:25169"
	},
	
	"Server": {
		"PortInfo": "0.0.0.0:13110"
	}
}
```
**Log.Level**: the logging level, it must be the one of **"error/warning/info/debug"**  
**Log.Suffix**: is the suffix of log file  
> ***"060102-15"*** means the name of logging file output is ending with **170328-20.log**.
**17** is short of year **2017**;
**03** is the month **March**;
**28** is the day of month;
**20** is the hour of the day. 

**Prog.CPU**: a limit of CPU usage. 0 is default, means to use all cores  
**Prog.Daemon**: you know  
**Prog.HealthPort**: the port for monitior, if you need  

**Server.PortInfo** the port for service 

Those items mentioned above are the base need of a server application. And they are defined in config file: ```gowork/app/conf/config.json```.

---
# Sample Building:
The sample code located in ```gowork/app/sample.go``` is an example of http server.

```
cd gowork/app
go build
```

> Ps: It will product an executable file name **app**.
   
---------------------------
### Fork me at GITHUB
There are more features, and shall we add them one by one?

```At last, Happy your coding.```
