package process

import (
	"time"

	"github.com/BeameryHQ/kubeaware/manager"
	"github.com/BeameryHQ/kubeaware/types"
	"github.com/MovieStoreGuy/artemis"
)

func init() {
	artemis.GetInstance().Log(artemis.Entry{artemis.Debug, "Adding process function into manager"})
	if err := manager.GetInstance().Register("process", New); err != nil {
		artemis.GetInstance().Log(artemis.Entry{artemis.Fatal, "Unable to load process due to: " + err.Error()})
	}
}

type mod struct {
	// Restarted keeps track of the number of times the process has been restarted
	// by the process manager
	restarted int           `monitor:"process.restart_count"`
	uptime    time.Duration `monitor:"process.up_time"`
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
		}
	}
	return nil
}
