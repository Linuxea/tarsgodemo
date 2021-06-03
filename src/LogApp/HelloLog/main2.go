package main



import (
"github.com/TarsCloud/TarsGo/tars"
"LogApp"
)

func main(){

	imp := new(SayHelloImp)
	app := new(LogApp.SayHello)
	cfg := tars.GetServerConfig()
	app.AddServant(imp, cfg.App + "." + cfg.Server + ".SayHelloObj")
	tars.Run()


}


