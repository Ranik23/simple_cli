package main

import (
	"cli/benchmarks/pprof_server"
	prom "cli/benchmarks/prometheus"
	"cli/config"
	cbr "cli/internal/cobra"
	read "cli/internal/commands/cat"
	"cli/internal/commands/count"
	"cli/internal/commands/count_strings"
	"cli/internal/metrics"
	"cli/internal/repository"
	"cli/internal/usecase"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	_ "github.com/prometheus/client_golang/prometheus/promhttp"
)

func main() {

	PrometheusServer := prom.NewPrometheusServer()
	PrometheusServer.MustRegister(metrics.MemoryUsage)
	prom.UpdateMemoryUsageGauge(metrics.MemoryUsage)
	PrometheusServer.StartPrometheusServer(":" + os.Getenv("PORT"))

	config, err := config.LoadConfig(os.Getenv("CONFIG"))
	if err != nil {
		log.Fatal("failed to load the config", err)
	}

	sigs := make(chan os.Signal, 1)
  	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)

	pprof_server := pprof_server.NewPprofServer("localhost:" + config.PprofServer.Host, config)
	pprof_server.Run()

	repo := repository.NewRepository()
	userCase := usecase.NewUserOperator(repo)

	cbr.RootCmd.AddCommand(read.CatCommand(userCase))
	cbr.RootCmd.AddCommand(count.CountCommand(userCase))
	cbr.RootCmd.AddCommand(count_strings.CountCommand2(userCase))

	if err := cbr.RootCmd.Execute(); err != nil {
		fmt.Println("Command execution error:", err)
		os.Exit(1)
	}

	if config.Env == "test" || config.Env == "docker" {
		<-sigs
	}
}
