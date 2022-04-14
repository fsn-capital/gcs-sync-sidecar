package main

import (
	"os"
	"os/exec"
	"os/signal"
	"syscall"
	"time"

	"github.com/spf13/viper"
)

func init() {
	// define env vars
	viper.BindEnv("SOURCE")
	viper.BindEnv("DESTINATION")
	viper.BindEnv("GOOGLE_APPLICATION_CREDENTIALS")
	// authenticate using service account
	cmd := exec.Command("gcloud", "auth", "activate-service-account", "--key-file", viper.GetString("GOOGLE_APPLICATION_CREDENTIALS"))
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err := cmd.Run()
	if err != nil {
		panic(err)
	}
}

func main() {
	// ticker setup
	ticker := time.NewTicker(time.Minute)
	// boilerplate for graceful shutdown
	done := make(chan bool)
	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)
	go func() {
		<-sigs
		ticker.Stop()
		done <- true
	}()
	// sync beforehand
	rsync()
	// main loop
	for {
		select {
		case <-done:
			return
		case <-ticker.C:
			rsync()
		}

	}
}

func rsync() {
	// use rsync to replicate changes
	cmd := exec.Command("gsutil", "-m", "rsync", "-r", "-d",
		viper.GetString("SOURCE"),
		viper.GetString("DESTINATION"),
	)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err := cmd.Run()
	if err != nil {
		panic(err)
	}
}
