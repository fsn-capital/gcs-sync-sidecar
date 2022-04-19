package common

import (
	"os"
	"os/exec"

	"github.com/spf13/viper"
)

func execute(name string, args ...string) error {
	cmd := exec.Command(name, args...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err := cmd.Run()
	return err
}

func Authenticate() error {
	return execute("gcloud", "auth", "activate-service-account", "--key-file", viper.GetString("GOOGLE_APPLICATION_CREDENTIALS"))
}

func Rsync() error {
	return execute("gsutil", "-m", "rsync", "-r", "-d", viper.GetString("SOURCE"), viper.GetString("DESTINATION"))
}
