package types

// Coordinator managers are the various parts of KubeAware
type Coordinator interface {
	// Configure will load a config and prepare the Coordinator
	// ready to be used later on.
	Configure(path string) error

	// Start causes the Coordinator to become aware of the events around it
	// and start running all of its sub processes
	Start() error
}
