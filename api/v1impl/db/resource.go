package db

import (
	"fmt"
	"time"

	"github.com/emicklei/go-restful"
	"github.com/juju/loggo"
)

var (
	logger = loggo.GetLogger("db")
)

const (
	newEndpointName = "new-endpoint-name"
)

// Resource is a resource DB access.
type Resource struct {
}

// NewResource creates new instance.
func NewResource() *Resource {
	records = make(map[string]string)
	return &Resource{}
}

// Register registers resource in restful container.
func (c *Resource) Register(container *restful.Container) *Resource {
	ws := new(restful.WebService)

	// MediaTypeApplicationYaml is a Mime Type for YAML
	const mediaTypeApplicationYaml = "application/x-yaml"

	ws.Path("/api"+"/db").
		Doc("DB endpoint provide get/add function").
		Consumes(restful.MIME_JSON, mediaTypeApplicationYaml).
		Produces(restful.MIME_JSON, mediaTypeApplicationYaml)

	ws.Route(ws.GET("").To(c.GetAllRecords).
		Doc("Get db record").
		Operation("Get db record"))

	ws.Route(ws.GET(fmt.Sprintf("{%s}", newEndpointName)).To(c.GetRecord).
		Doc("Get db record").
		Operation("Get db record"))

	ws.Route(ws.POST(fmt.Sprintf("{%s}", newEndpointName)).To(c.CreateRecord).
		Doc("Add new record").
		Operation("Add new record"))

	container.Add(ws)

	return c
}

type Record struct {
	Endpoint  string `json:"endpoint" yaml:"endpoint"`
	Timestamp string `json:"timestamp" yaml:"timestamp"`
}

var records map[string]string

// GetAllRecords gets all record
func (c *Resource) GetAllRecords(request *restful.Request, response *restful.Response) {
	response.WriteEntity(records)
}

// GetRecord gets record
func (c *Resource) GetRecord(request *restful.Request, response *restful.Response) {
	name := request.PathParameter(newEndpointName)
	response.WriteEntity(records[name])
}

// CreateRecord creates new record
func (c *Resource) CreateRecord(request *restful.Request, response *restful.Response) {
	name := request.PathParameter(newEndpointName)

	records[name] = time.Now().String()

	response.WriteEntity("success")
}
