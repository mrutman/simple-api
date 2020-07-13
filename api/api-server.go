package api

import (
	"bytes"
	"fmt"
	"log"
	"net/http"
	"os"
	"path"
	"runtime"

	"github.com/mrutman/simple-api/api/v1impl"
	"github.com/mrutman/simple-api/pkg/config"

	"github.com/emicklei/go-restful"
	"github.com/gorilla/handlers"
	"github.com/juju/loggo"
)

// Server is HTTP Server for API.
type Server struct {
	api *v1impl.SimpleAPI
}

// logger for API server.
var logger = loggo.GetLogger("API")

// NewServer creates a new Kublr API server but does not configure it.
// Call RegisterAndServe to register REST endpoints and start serving.
func NewServer(api *v1impl.SimpleAPI) *Server {
	server := &Server{
		api: api,
	}
	return server
}

// RegisterAndServe registers REST endpoints and starts serving HTTP server.
func (apiServer *Server) RegisterAndServe() error {
	//restful.EnableTracing(true)
	wsContainer := restful.NewContainer()
	wsContainer.RecoverHandler(recoveryHandler)
	wsContainer.DoNotRecover(false)

	err := apiServer.api.Register(wsContainer, true)
	if err != nil {
		return err
	}

	return apiServer.serve(wsContainer)
}

// serve starts HTTP serving.
func (apiServer *Server) serve(wsContainer *restful.Container) error {
	cfg := config.GetSimpleAPIConfig()
	log.Printf("start listening on localhost:%s", cfg.ServerPort)

	server := &http.Server{
		Addr:    ":" + cfg.ServerPort,
		Handler: handlers.LoggingHandler(os.Stdout, wsContainer),
	}

	return server.ListenAndServe()
}

// recoveryHandler catches panics and logs them.
// Returns 500 Error to the caller.
func recoveryHandler(panicReason interface{}, httpWriter http.ResponseWriter) {
	logger.Errorf("[restful] recover from panic situation: - %v\r\n", panicReason)
	var buffer bytes.Buffer
	buffer.WriteString(fmt.Sprintf("stack trace:\n"))
	for i := 2; ; i++ {
		_, file, line, ok := runtime.Caller(i)
		if !ok {
			break
		}
		buffer.WriteString(fmt.Sprintf("    %s:%d\r\n", path.Base(file), line))
	}
	logger.Debugf(buffer.String())
	httpWriter.WriteHeader(http.StatusInternalServerError)
}
