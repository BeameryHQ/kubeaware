package types

// Coordinator managers are the various parts of KubeAware
type Coordinator interface {
	// Configure will load a config and prepare the Coordinator
	// ready to be used later on.
	Configure(path string) error

	// RunProcess will start a subprocess and manage it.
	// Any updates within kubernetes can cause the process to restart
	// RunProcess(path string, args ...string) error

	// Start causes the Coordinator to become aware of the events around it
	// and start running all of its sub processes
	Start() error
}
