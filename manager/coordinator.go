package manager

import (
	"fmt"
	"os"
	"os/signal"
	"sync"
	"syscall"

	"github.com/BeameryHQ/kubeaware/types"
	"github.com/MovieStoreGuy/artemis"
)

var (
	once      sync.Once
	kubeAware *mon
)

type mon struct {
	modules map[string]func() types.Module
	loaded  []types.Module
	mutex   sync.Mutex
}

// GetInstance returns the signalton that is used within the main program
func GetInstance() types.Coordinator {
	once.Do(func() {
		artemis.GetInstance().Log(artemis.InfoEntry("Creating manager instance"))
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
	return m.loaded
}

func (m *mon) Register(name string, mod func() types.Module) error {
	// As we can not ensure that init functions happen synchronisly
	// then a lock needs to be introduced.
	m.mutex.Lock()
	defer m.mutex.Unlock()
	artemis.GetInstance().Log(artemis.InfoEntry("Attempting to register new module:", name))
	if _, exist := m.modules[name]; exist {
		return fmt.Errorf("The module |%v| already exists", name)
	}
	m.modules[name] = mod
	return nil
}

func (m *mon) Start() error {
	// Load all the modules ready for monitoring.
	for _, mod := range m.loaded {
		if err := exportVariables(mod); err != nil {
			artemis.GetInstance().Log(artemis.FatalEntry("Unable to load module due to: ", err))
		}
		go mod.Start()
	}
	artemis.GetInstance().Log(artemis.InfoEntry("Now running all modules"))
	return m.awaitSignals()
}

// awaitSignals will listen for predefined signals and will
func (m *mon) awaitSignals() error {
	artemis.GetInstance().Log(artemis.InfoEntry("Attaching signal handlers"))
	done := make(chan error)
	sigs := make(chan os.Signal, 1)
	go func() {
		event := <-sigs
		artemis.GetInstance().Log(artemis.InfoEntry("recieved event: ", event))
		for _, module := range m.loaded {
			switch event {
			case syscall.SIGINT:
				module.ExitWithCondition(types.Shutdown)
			case syscall.SIGABRT:
				module.ExitWithCondition(types.ForceShutdown)
			}
		}
		done <- nil
	}()
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGABRT)
	return <-done
}
