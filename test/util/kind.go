package util

import (
	"os"
	"os/exec"
)

func LoadImageToKindClusterWithName(names ...string) error {
	cluster := "kvbmc-e2e"
	if v, ok := os.LookupEnv("KIND_CLUSTER"); ok {
		cluster = v
	}
	kindOptions := append([]string{"load", "docker-image", "--name", cluster}, names...)
	cmd := exec.Command("kind", kindOptions...)
	_, err := Run(cmd)
	return err
}
