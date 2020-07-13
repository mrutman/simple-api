package config

import (
	"github.com/mrutman/simple-api/pkg/config"

	"github.com/emicklei/go-restful"
)

// Resource is a resource for config.
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

	ws.Path("/api"+"/config").
		Doc("Config endpoint returns current simple-api config").
		Consumes(restful.MIME_JSON, mediaTypeApplicationYaml).
		Produces(restful.MIME_JSON, mediaTypeApplicationYaml)

	ws.Route(ws.GET("").To(c.GetConfig).
		Doc("Returns current config").
		Operation("operation returns current config"))

	container.Add(ws)

	return c
}

// GetConfig returns current config
func (c *Resource) GetConfig(request *restful.Request, response *restful.Response) {
	response.WriteEntity(config.GetSimpleAPIConfig())
}
