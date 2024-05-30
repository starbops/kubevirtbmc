package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"

	"github.com/sirupsen/logrus"
	"github.com/urfave/cli/v2"

	"kubevirt.org/virtualmachinebmc/pkg/kbmc"
)

func main() {
	var options kbmc.Options

	app := &cli.App{
		Name:  "kbmc",
		Usage: "receive ipmi requests and traslate them into native k8s api calls",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:        "address",
				Aliases:     []string{"a"},
				Value:       "127.0.0.1",
				Usage:       "listen on `IP ADDRESS`",
				Destination: &options.Address,
			},
			&cli.StringFlag{
				Name:        "kubeconfig",
				Aliases:     []string{"k"},
				Value:       "/root/.kube/config",
				Usage:       "use a specific kubeconfig `FILE`",
				Destination: &options.KubeconfigPath,
			},
			&cli.IntFlag{
				Name:        "port",
				Aliases:     []string{"p"},
				Value:       10623,
				Usage:       "listen on `PORT`",
				Destination: &options.Port,
			},
		},
		Action: func(cCtx *cli.Context) error {
			ctx := context.WithValue(cCtx.Context, "VMNamespace", cCtx.Args().Get(0))
			ctx = context.WithValue(ctx, "VMName", cCtx.Args().Get(1))
			return run(ctx, options)
		},
	}

	if err := app.Run(os.Args); err != nil {
		logrus.Fatal(err)
	}
}

func run(ctx context.Context, options kbmc.Options) error {
	logrus.Info("Starting kBMC")

	// TODO: check kubeconfig flag instead
	// check whether we're in a cluster or not
	_, ok := os.LookupEnv("KUBERNETES_SERVICE_HOST")

	// trap Ctrl+C can call cancel on the context
	ctx, cancel := context.WithCancel(ctx)
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	defer func() {
		signal.Stop(c)
		cancel()
	}()
	go func() {
		select {
		case <-c:
			cancel()
		case <-ctx.Done():
		}
		<-c
		os.Exit(2)
	}()

	kbmc, err := kbmc.NewKBMC(ctx, options, ok)
	if err != nil {
		return fmt.Errorf("failed to create kbmc server: %v", err)
	}

	return kbmc.Run()
}
