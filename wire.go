package main

import "github.com/google/wire"

func InitializeEvent(msg string) Event {
	wire.Build(NewEvent, NewGreeter, NewMessage)
	return Event{} //返回值没有实际意义，只需符合函数签名即可
}
