package main

import (
	_case "CloudNative/sync/case"
	"context"
	"os"
	"os/signal"
)

func main() {
	//_case.WaitGroup()
	//_case.WaitGroupCase1()
	//_case.CondCase()
	_case.CondQueueCase()
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, os.Kill)
	defer stop()
	<-ctx.Done()
}
