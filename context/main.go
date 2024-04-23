package main

import (
	_case "CloudNative/context/case"
	"context"
	"os"
	"os/signal"
)

func main() {
	_case.ContextCase()

	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, os.Kill)
	defer stop()
	<-ctx.Done()
}
