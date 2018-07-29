package manager

import (
	"fmt"
	"sync"

	"github.com/BeameryHQ/kubeaware/types"
)

var (
	once      sync.Once
	kubeAware *mon
)

type mon struct {
	modules map[string]func() types.Module
	mutex   sync.Mutex
}

// GetInstance returns the signalton that is used within the main program
func GetInstance() types.Coordinator {
	once.Do(func() {
		kubeAware = &mon{}
	})
	return kubeAware
}

func (m *mon) Configure(path string) error {
	// Once we have loaded the config as a byte buffer
	// Then we are going to itterate over each loaded module and pass
	// the buff as one of the args
	return nil
}

func (m *mon) LoadedModules() []types.Module {
	return []types.Module{}
}

func (m *mon) Register(name string, mod func() types.Module) error {
	// As we can not ensure that init functions happen synchronisly
	// then a lock needs to be introduced.
	m.mutex.Lock()
	defer m.mutex.Unlock()
	if _, exist := m.modules[name]; exist {
		return fmt.Errorf("The module |%v| already exists", name)
	}
	m.modules[name] = mod
	return nil
}

func (m *mon) Start() error {
	return nil
}
