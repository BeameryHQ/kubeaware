package process

import (
	"time"

	"github.com/BeameryHQ/kubeaware/types"
)

type mod struct {
	// Restarted keeps track of the number of times the process has been restarted
	// by the process manager
	restarted int           `monitor:"process.restart_count"`
	uptime    time.Duration `monitor:"process.up_time"`
	process   inner         `monitor:"-"`
}

type inner struct {
	metrics float32 `monitor:"process.inner.metrics"`
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
	start := time.Now()
	ticker := time.NewTicker(time.Second)
	for {
		select {
		case <-ticker.C:
			m.restarted++
			m.uptime = time.Since(start)
			m.process.metrics = 3.14
		}
	}
	return nil
}
