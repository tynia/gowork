# [gowork](https://github.com/tynia/gowork.git)
**GOWORK** is a tiny server framework implement using golang.  
It is aim at developing server faster.  
It is also an exercise to improve myself.  

# Author
**name**  : **tynia**  
**email** : [ekesi@live.com](ekesi@live.com)

# LICENSE
The **gowork** library is released under the [**MIT LICENSE**](http://opensource.org/licenses/mit-license.php).  
you can use it without any copyright about it.  

# Besides
I'd like you will leave my name(tynia) and [repository](https://github.com/tynia/gowork) in your project, which you use the [**gowork**](https://github.com/tynia/gowork) in it.  

---
The **gowork** framework contains several modules includeing:
- application 
- [logger](https://github.com/xuyu/logging)   --[**@xuyu**](https://github.com/xuyu) 

# Introduction
**[application]** 

It is a simple application .  
When using the logger, all you need do like this below:
```
import (
    "gowork/common/application"
    "gowork/common/application"
    ... // other imports you need
)

...

func main() {
    config := map[string]interface{}{}
    value := "hello, gowork !"
    config["key"] = value
    app := application.NewApplication("sample", config)
    v := app.Get("key")
    str, _ := convertor.ToString(v) // you can get value of "key" is "hello, gowork !"

    err := app.Go()
    if err != nil {
	    fmt.Println("Error: %s", err.Error())
    }
}
```

Notice: In application.Go(), it will parse a json file. the file **must be** like this: 
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

This is the base need of a server application. These are defined in config file: ```gowork/app/conf/config.json``` 

**Log.Level**: the logging level, it must be the one of **"error/warning/info/debug"**
**Log.Suffix**: is the suffix of log file, 
> ***"060102-15"*** means the name of logging file output is ending with 170328-20. 17 is short of year 2017; 03 is the month; 28 is the day of month; 20 is the hour of the day. 

**Prog.CPU**: a limit of CPU usage. 0 is default, means to use all cores 
**Prog.Daemon**: you know
**Prog.HealthPort**: the port for monitior, if you need 

**Server.PortInfo** the port for service 

More detail, please see ```gowork/app/sample.go```

 
---
# Sample Building:
The sample code located in app/sample.go is a simple example of http server.

```
cd gowork/app
go build
```

> Ps: It will product an executable file name **app**.
   
---------------------------
### Fork me at GITHUB
There are more features, and shall we add them one by one?

```At last, Happy your coding.```