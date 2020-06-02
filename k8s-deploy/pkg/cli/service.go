package cli

import (
	"github.com/FederatedAI/KubeFATE/k8s-deploy/pkg/api"
	"github.com/urfave/cli/v2"
)

func serviceCommand() *cli.Command {
	return &cli.Command{
		Name:   "service",
		Usage:  "Start KubeFATE as service",
		Action: serviceRun,
	}
}

func serviceRun(c *cli.Context) error {
	api.Run()
	return nil
}
