package main

import (
	log "github.com/sirupsen/logrus"
	"gitlab.com/ConsenSys/client/fr/core-stack/service/contract-registry.git/cmd"
)

func main() {
	command := cmd.NewCommand()
	if err := command.Execute(); err != nil {
		log.WithError(err).Fatalf("main: execution failed")
	}
	log.Infof("main: execution completed")
}
