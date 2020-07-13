package db

import (
	"fmt"
	"net/http"

	"github.com/mrutman/simple-api/pkg/db"

	"github.com/emicklei/go-restful"
	"github.com/juju/loggo"
)

var (
	logger = loggo.GetLogger("ep-db")
)

const (
	newEndpointName = "new-endpoint-name"
)

// Resource is a resource DB access.
type Resource struct {
}

// NewResource creates new instance.
func NewResource() *Resource {
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

// GetAllRecords gets all record
func (c *Resource) GetAllRecords(request *restful.Request, response *restful.Response) {
	all, err := db.GetAllSimpleRecords()
	if err != nil {
		err = response.WriteHeaderAndEntity(http.StatusInternalServerError,
			fmt.Sprintf("Failed to get all records: '%v'", err))
		if err != nil {
			logger.Errorf("Error: '%v'", err)
		}
		return
	}

	response.WriteEntity(all)
}

// GetRecord gets record
func (c *Resource) GetRecord(request *restful.Request, response *restful.Response) {
	name := request.PathParameter(newEndpointName)
	simpleRecord, err := db.GetSimpleRecord(name)
	if err != nil {
		err = response.WriteHeaderAndEntity(http.StatusInternalServerError,
			fmt.Sprintf("Failed to get record for '%s': '%v'", name, err))
		if err != nil {
			logger.Errorf("Error: '%v'", err)
		}
		return
	}
	response.WriteEntity(simpleRecord)
}

// CreateRecord creates new record
func (c *Resource) CreateRecord(request *restful.Request, response *restful.Response) {
	c.GetRecord(request, response)
}
