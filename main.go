package main

import (
	"github.com/koki/short-drone-plugin/cmd"

	log "github.com/Sirupsen/logrus"
)

func main() {
	if err := cmd.KokiCmd.Execute(); err != nil {
		log.Fatal(err)
	}
}
