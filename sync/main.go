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
	//_case.CondQueueCase()
	//_case.MutexCase()
	//_case.MapCase()
	//_case.MapCase1()
	//_case.PoolCase()
	//_case.OnceCase()
	//_case.AtomicCase()
	//_case.AtomicCase1()
	_case.AtomicCase2()
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, os.Kill)
	defer stop()
	<-ctx.Done()
}
