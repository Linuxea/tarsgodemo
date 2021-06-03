package main

import (
	"context"
)

// SayHelloObjImp servant implementation
type SayHelloObjImp struct {
}

// Init servant init
func (imp *SayHelloObjImp) Init() (error) {
	//initialize servant here:
	//...
	return nil
}

// Destroy servant destory
func (imp *SayHelloObjImp) Destroy() {
	//destroy servant here:
	//...
}

func (imp *SayHelloObjImp) Add(ctx context.Context, a int32, b int32, c *int32) (int32, error) {
	//Doing something in your function
	//...
	return 0, nil
}
func (imp *SayHelloObjImp) Sub(ctx context.Context, a int32, b int32, c *int32) (int32, error) {
	//Doing something in your function
	//...
	return 0, nil
}
