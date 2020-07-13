package cmd

import (
	"github.com/mrutman/simple-api/api"
	"github.com/mrutman/simple-api/api/v1impl"

	"github.com/juju/loggo"
)

var logger = loggo.GetLogger("cmd")

func Run() {
	logger.SetLogLevel(loggo.INFO)

	simpleAPI := v1impl.NewSimpleAPI()

	simpleServer := api.NewServer(simpleAPI)
	if err := simpleServer.RegisterAndServe(); err != nil {
		panic(err)
	}
}
