package main

import (
	"flag"
	"os"

	"github.com/BeameryHQ/kubeaware/manager"
	"github.com/MovieStoreGuy/artemis"
)

var (
	confPath = os.Getenv("CONFIG_PATH")
)

func init() {
	artemis.GetInstance().Set(artemis.Info, os.Stdout)
	const (
		blank = ""
	)
	flag.StringVar(&confPath, "configure", blank, "defines the path to load the config from (ENV: CONFIG_PATH)")
}

func main() {
	// Ensure all log messages happen
	defer artemis.GetInstance().Stop()
	flag.Parse()
	// Load the config for KubeAware
	if err := manager.GetInstance().Configure(confPath); err != nil {
		artemis.GetInstance().Log(artemis.Entry{artemis.Fatal, "Unable to configure instance due to:" + err.Error()})
	}
	// Determine which kubernetes types we need to watch in order to become "Aware"
	if err := manager.GetInstance().Start(); err != nil {
		artemis.GetInstance().Log(artemis.Entry{artemis.Fatal, "Unable to start instance due to:" + err.Error()})
	}
}
