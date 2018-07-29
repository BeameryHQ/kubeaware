package types

// Module is a simple unit of defined functionality that is used within
// kubeAware
type Module interface {
	// Connect(cord Coordinator) error

	ExitWithCondition(cond Condition)

	// ParseConfig takes a loaded yaml file as a bytes array
	// ready to marshal the content into the module if it requires it.
	ParseConfig(info []byte) error

	// Start simply starts the module
	Start() error
}
