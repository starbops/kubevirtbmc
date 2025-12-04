package redfish

import (
	"context"
	"fmt"
	"net/http"
	"sync"

	"github.com/sirupsen/logrus"
	"kubevirt.io/kubevirtbmc/pkg/generated/redfish/server"
	"kubevirt.io/kubevirtbmc/pkg/resourcemanager"
	"kubevirt.io/kubevirtbmc/pkg/session"
)

type Emulator struct {
	ctx  context.Context
	port int

	bmcUser     string
	bmcPassword string

	wg     sync.WaitGroup
	server *http.Server
}

func NewEmulator(ctx context.Context, port int, bmcUser string, bmcPassword string, resourceManager resourcemanager.ResourceManager) *Emulator {
	apiService := NewAPIService(bmcUser, bmcPassword, resourceManager)
	apiController := server.NewDefaultAPIController(apiService)
	router := server.NewRouter(session.AuthMiddleware(bmcUser, bmcPassword), apiController)

	return &Emulator{
		ctx:         ctx,
		port:        port,
		bmcUser:     bmcUser,
		bmcPassword: bmcPassword,
		server: &http.Server{
			Addr:    fmt.Sprintf(":%d", port),
			Handler: router,
		},
	}
}

func (e *Emulator) Run() error {
	e.wg.Add(1)

	go func() {
		defer e.wg.Done()

		if err := e.server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			fmt.Println(err)
		}
	}()

	return nil
}

func (e *Emulator) Stop() {
	if err := e.server.Shutdown(e.ctx); err != nil {
		fmt.Println(err)
	}
	e.wg.Wait()
	logrus.Info("Redfish emulator gracefully stopped")
}
