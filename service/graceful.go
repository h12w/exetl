package service

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"
)

// NotifyStop listens to process signal and calls stopFn when received
func NotifyStop(stopFn func(context.Context) error) {
	stopChan := make(chan os.Signal)
	signal.Notify(stopChan, syscall.SIGTERM, syscall.SIGINT)
	go func() {
		sig := <-stopChan
		log.Printf("got signal %v", sig)
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()
		if err := stopFn(ctx); err != nil {
			log.Print(err)
		}
	}()
}
