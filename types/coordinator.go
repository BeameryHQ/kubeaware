package types

// Coordinator managers are the various parts of KubeAware
type Coordinator interface {
	// Configure will load a config and prepare the Coordinator
	// ready to be used later on.
	Configure(path string) error

	// Register will load the given module and make it ready for
	// for the user at runtime.
	// This function should only be used to in each module's init function.
	Register(name string, loader func() Module) error

	// LoadedModules returns all the currently loaded modules within the Coordinator
	LoadedModules() []Module

	// Start causes the Coordinator to become aware of the events around it
	// and start running all of its sub processes
	Start() error
}
