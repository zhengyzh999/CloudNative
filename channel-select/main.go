package main

import (
	_case "CloudNative/channel-select/case"
	"os"
	"os/signal"
)

func main() {
	//_case.Communication()
	//_case.ConcurrentSync()
	_case.NoticeAndMultipleExiting()

	// 利用阻塞机制
	ch := make(chan os.Signal)
	signal.Notify(ch, os.Interrupt, os.Kill)
	<-ch
}
