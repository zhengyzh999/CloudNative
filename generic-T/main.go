package main

import (
	"context"
	_case "go-base/generic-T/case"
	"os"
	"os/signal"
)

func main() {
	_case.SimpleCase()
	_case.CusNumTCase()
	_case.BuildInCase()
	_case.TTypeCase()
	_case.TTypeCase1()
	_case.InterfaceCase()
	_case.ReceiverCase()

	// 不要让主协程太快退出
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, os.Kill)
	defer stop()
	<-ctx.Done()
}
