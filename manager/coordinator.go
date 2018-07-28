package manager

import (
	"sync"

	"github.com/BeameryHQ/kubeaware/types"
)

var (
	once      sync.Once
	kubeAware *mon
)

type mon struct {
}

// GetInstance returns the signalton that is used within the main program
func GetInstance() types.Coordinator {
	once.Do(func() {
		kubeAware = &mon{}
	})
	return kubeAware
}

func (m *mon) Configure(path string) error {

	return nil
}

func (m *mon) Start() error {
	return nil
}
