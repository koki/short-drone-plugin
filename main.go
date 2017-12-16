package main

import (
	"github.com/kubeciio/koki/cmd"

	log "github.com/Sirupsen/logrus"
)

func main() {
	if err := cmd.KokiCmd.Execute(); err != nil {
		log.Fatal(err)
	}
}
