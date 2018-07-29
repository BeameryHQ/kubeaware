package process

import (
	"github.com/BeameryHQ/kubeaware/manager"
	"github.com/BeameryHQ/kubeaware/types"
)

func init() {
	if err := manager.GetInstance().Register("process", New); err != nil {
		panic(err)
	}
}

type mod struct {
	// Restarted keeps track of the number of times the process has been restarted
	// by the process manager
	restarted int `monitor:"process.restart_count"`
}

func New() types.Module {
	return &mod{}
}

func (m *mod) ExitWithCondition(cond types.Condition) {
	// Do something based on the events
}

func (m *mod) ParseConfig(info []byte) error {
	return nil
}

func (m *mod) Start() error {
	return nil
}
