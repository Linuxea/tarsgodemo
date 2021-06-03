package main

import (
	"LogApp"
	"fmt"
	"github.com/TarsCloud/TarsGo/tars"
)

var comm *tars.Communicator

func main() {

	comm = tars.NewCommunicator()
	obj := "LogApp.HelloLog.SayHelloObj"
	// tarsregistry service at 192.168.1.1:17890
	comm.SetProperty("locator", "tars.tarsregistry.QueryObj@tcp -h 172.25.0.3 -p 17890")
	app := new(LogApp.SayHello)

	comm.StringToProxy(obj, app)
	reqStr := "tars"
	var resp string
	ret, err := app.EchoHello(reqStr, &resp)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("hhhhhhhhhhhhhhh ret:", ret, "resp:", resp)

}
