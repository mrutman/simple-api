package v1impl

import (
	"log"
	"os"

	"github.com/emicklei/go-restful"
	"github.com/juju/loggo"

	"github.com/mrutman/simple-api/api/v1impl/config"
	"github.com/mrutman/simple-api/api/v1impl/db"
)

var logger = loggo.GetLogger("SimpleAPI")

// SimpleAPI is a definition of Simple API.
type SimpleAPI struct {
}

// NewSimpleAPI creates new instance of Simple API.
// It is required to call Register before start to use it.
func NewSimpleAPI() *SimpleAPI {

	api := &SimpleAPI{}

	restful.DefaultRequestContentType(restful.MIME_JSON)
	restful.DefaultResponseContentType(restful.MIME_JSON)
	restful.SetLogger(log.New(os.Stderr, "", log.LstdFlags|log.Lshortfile|log.Lmicroseconds))

	return api
}

// Register registers REST resources in container.
func (api *SimpleAPI) Register(wsContainer *restful.Container, insecure bool) error {
	config.NewResource().Register(wsContainer)
	db.NewResource().Register(wsContainer)
	return nil
}
