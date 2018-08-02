package main

import (
	"flag"

	"github.com/BeameryHQ/kubeaware/manager"
)

func main() {
	flag.Parse()
	// Load the config for KubeAware
	if err := manager.GetInstance().Configure(""); err != nil {
		panic(err)
	}
	// Determine which kubernetes types we need to watch in order to become "Aware"
	if err := manager.GetInstance().Start(); err != nil {
		panic(err)
	}
	// Loop until we have been told to shutdown
}
