package main



type SayHelloImp struct {}



func (imp *SayHelloImp) EchoHello(name string, greeting *string)(int32, error) {
    *greeting = "hello my boy 你好啊 " + name
    return 0, nil

}


