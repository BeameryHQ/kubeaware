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
