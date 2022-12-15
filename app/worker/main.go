package main

import (
	"fmt"

	"github.com/nndd91/cadence-api-example/app/adapters/cadenceAdapter"
	"github.com/nndd91/cadence-api-example/app/config"
	"github.com/nndd91/cadence-api-example/app/worker/workflows"
	"go.uber.org/cadence/worker"
	"go.uber.org/zap"
)

func startWorker(h *cadenceAdapter.CadenceAdapter, taskList string) {
	workerOptions := worker.Options{
		MetricsScope: h.Scope,
		Logger:       h.Logger,
	 }
  
	 cadenceWorker := worker.New(h.ServiceClient, h.Config.Domain, taskList, workerOptions)
	 err := cadenceWorker.Start()
	 if err != nil {
		h.Logger.Error("Failed to start workers.", zap.Error(err))
		panic("Failed to start workers")
	 }
}

func main() {
	fmt.Println("Starting worker...")
	var appConfig config.AppConfig
	appConfig.Setup()
	var cadenceClient cadenceAdapter.CadenceAdapter
	cadenceClient.Setup(&appConfig.Cadence)
	startWorker(&cadenceClient, workflows.TaskListName)
	select{}
}