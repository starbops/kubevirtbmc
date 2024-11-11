package util

import (
	"os"
	"os/exec"
)

func LoadImageToKindClusterWithName(name string) error {
	cluster := "kvbmc-e2e"
	if v, ok := os.LookupEnv("KIND_CLUSTER"); ok {
		cluster = v
	}
	kindOptions := []string{"load", "docker-image", name, "--name", cluster}
	cmd := exec.Command("kind", kindOptions...)
	_, err := Run(cmd)
	return err
}
