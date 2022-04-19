package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/fsn-capital/gcs-sync-sidecar/common"
	"github.com/spf13/viper"
)

func init() {
	// define env vars
	viper.BindEnv("SOURCE")
	viper.BindEnv("DESTINATION")
	viper.BindEnv("GOOGLE_APPLICATION_CREDENTIALS")
	// authenticate using service account
	err := common.Authenticate()
	if err != nil {
		errLogger("authenticate", err)
	}
}

func main() {
	// ticker setup
	ticker := time.NewTicker(time.Minute)
	// boilerplate code for graceful shutdown
	done := make(chan bool)
	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)
	go func() {
		<-sigs
		ticker.Stop()
		done <- true
	}()
	// sync beforehand
	err := common.Rsync()
	if err != nil {
		errLogger("rsync", err)
	}
	// main loop
	for {
		select {
		case <-done:
			return
		case <-ticker.C:
			err := common.Rsync()
			if err != nil {
				errLogger("rsync", err)
			}
		}

	}
}

func errLogger(name string, err error) {
	log.Fatalf("Observed error during %s: %s", name, err.Error())
}
