package server

import (
	"expvar"
	"net/http"
	"time"

	"github.com/BeameryHQ/kubeaware/manager"
	"github.com/BeameryHQ/kubeaware/types"
	"github.com/MovieStoreGuy/artemis"
	"github.com/gorilla/mux"
)

type module struct {
	server *http.Server
	router *mux.Router
}

func init() {
	if err := manager.GetInstance().Register("server", New); err != nil {
		artemis.GetInstance().Log(artemis.FatalEntry("Unable to register server due to: ", err))
	}
}

func New() types.Module {
	return &module{
		server: &http.Server{
			ReadTimeout:    10 * time.Second,
			WriteTimeout:   10 * time.Second,
			MaxHeaderBytes: 1 << 20,
		},
		router: mux.NewRouter(),
	}
}

func (m *module) ExitWithCondition(cond types.Condition) {
	switch cond {
	case types.ForceShutdown, types.Shutdown:
		if err := m.server.Close(); err != nil {
			artemis.GetInstance().Log(artemis.Entry{artemis.Info, "Server failed to shutdown due to: " + err.Error()})
		}
	}
}

func (m *module) ParseConfig(buff []byte) error {
	// Hard coded for now until I decide on the ParseConfig
	m.server.Addr = ":8000"
	m.server.Handler = m.router
	return nil
}

func (m *module) Start() error {
	m.enableDebugExport()
	return m.server.ListenAndServe()
}

func (m *module) enableDebugExport() {
	m.router.Handle("/debug/vars", expvar.Handler())
}
