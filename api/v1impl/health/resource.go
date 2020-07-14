package health

import (
	"fmt"
	"net/http"

	"github.com/mrutman/simple-api/pkg/db"

	"github.com/emicklei/go-restful"
	"github.com/juju/loggo"
)

var (
	logger = loggo.GetLogger("health")
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

// Register performs registration of REST resource in the container.
func (h *Resource) Register(container *restful.Container) *Resource {
	// Add healthz endpoints
	container.Add(h.createCheckEndpoint("/healthz", h.HealthCheck,
		"Health check", "Perform health check", "Returns Code 200 if service is up and running"))
	container.Add(h.createCheckEndpoint("/api/healthz", h.HealthCheck,
		"Health check", "Perform health check", "Returns Code 200 if service is up and running"))

	// Add readyz endpoints
	container.Add(h.createCheckEndpoint("/readyz", h.ReadinessCheck,
		"Readiness check", "Perform readiness check", "Returns Code 200 if service is up and running and database is accessible"))
	container.Add(h.createCheckEndpoint("/api/readyz", h.ReadinessCheck,
		"Readiness check", "Perform readiness check", "Returns Code 200 if service is up and running and database is accessible"))

	// Add livez endpoints
	container.Add(h.createCheckEndpoint("/livez", h.LivenessCheck,
		"Liveness check", "Perform liveness check", "Returns Code 200 if service is up and running"))
	container.Add(h.createCheckEndpoint("/api/livez", h.LivenessCheck,
		"Liveness check", "Perform liveness check", "Returns Code 200 if service is up and running"))

	return h
}

func (h *Resource) createCheckEndpoint(path string, do restful.RouteFunction, pathDoc, routeDoc, retDoc string) *restful.WebService {
	ws := new(restful.WebService)
	ws.Path(path).
		Doc(pathDoc).
		Produces(restful.MIME_JSON)
	ws.Route(ws.GET("").To(do).
		Doc(routeDoc).
		Returns(http.StatusOK, retDoc, nil))

	return ws
}

// HealthCheck checks service is alive
func (h *Resource) HealthCheck(request *restful.Request, response *restful.Response) {
	response.WriteHeaderAndEntity(http.StatusOK, "ok")
}

// ReadinessCheck checks service is ready
func (h *Resource) ReadinessCheck(request *restful.Request, response *restful.Response) {
	_, err := db.GetAllSimpleRecords()
	if err != nil {
		err = response.WriteHeaderAndEntity(http.StatusInternalServerError,
			fmt.Sprintf("Database is not ready: '%v'", err))
		if err != nil {
			logger.Errorf("Error: '%v'", err)
		}
		return
	}

	response.WriteHeaderAndEntity(http.StatusOK, "ok")
}

// LivenessCheck checks service is ready
func (h *Resource) LivenessCheck(request *restful.Request, response *restful.Response) {
	response.WriteHeaderAndEntity(http.StatusOK, "ok")
}
