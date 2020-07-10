package cmd

import (
	"fmt"

	"github.com/juju/loggo"

	"github.com/mrutman/simple-api/api"
	"github.com/mrutman/simple-api/api/v1impl"
)

var logger = loggo.GetLogger("cmd")

func Run() {
	logger.SetLogLevel(loggo.INFO)
	fmt.Printf("start\n")

	logger.Errorf("run")
	logger.Infof("run")

	simpleAPI := v1impl.NewSimpleAPI()

	simpleServer := api.NewServer(simpleAPI)
	if err := simpleServer.RegisterAndServe(); err != nil {
		panic(err)
	}
}
