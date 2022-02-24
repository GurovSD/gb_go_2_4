package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	sigs := make(chan os.Signal, 1)

	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		//ожидание сигнала в канале sigs
		sig := <-sigs
		fmt.Println(sig)
	}()

}
