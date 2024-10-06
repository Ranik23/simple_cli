package main

import (
	"cli/benchmarks/pprof_server"
	prom "cli/benchmarks/prometheus"
	"cli/config"
	cbr "cli/internal/cobra"
	read "cli/internal/commands/cat"
	"cli/internal/commands/count"
	"cli/internal/commands/count_strings"
	"cli/internal/commands/ls"
	"cli/internal/commands/clear"
	"cli/internal/metrics"
	"cli/internal/repository"
	"cli/internal/usecase"
	"fmt"
	"log"
	"os"
	"os/signal"
	"path/filepath"
	"syscall"

	_ "github.com/prometheus/client_golang/prometheus/promhttp"
)

var (
	ErrNoHomeDir = fmt.Errorf("failed to get home dir")
)

func main() {

	PrometheusServer := prom.NewPrometheusServer()
	PrometheusServer.MustRegister(metrics.MemoryUsage)
	prom.UpdateMemoryUsageGauge(metrics.MemoryUsage)
	PrometheusServer.StartPrometheusServer(":" + os.Getenv("PORT"))

	homeDir, err := os.UserHomeDir()
	if err != nil {
		fmt.Println("failed to get home dir")
		return 
	} 

	configPath := filepath.Join(homeDir, ".config", "cli", "config.yml")

	config, err := config.LoadConfig(configPath)
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
	cbr.RootCmd.AddCommand(ls.LsCommand(*userCase))
	cbr.RootCmd.AddCommand(clear.ClearCommand(*&userCase))

	if err := cbr.RootCmd.Execute(); err != nil {
		fmt.Println("Command execution error:", err)
		os.Exit(1)
	}

	if config.Env == "test" || config.Env == "docker" {
		<-sigs
	}
}
