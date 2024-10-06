package prometheus

import (
	"net/http"
	"runtime"
	"time"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)


type PrometheusServer struct {}

func NewPrometheusServer() *PrometheusServer {
	return &PrometheusServer{}
}

func (p *PrometheusServer) MustRegister(metrics ...prometheus.Collector) {
	prometheus.MustRegister(metrics...)
}

func (p *PrometheusServer) StartPrometheusServer(addr string) {
	go func() {
		http.Handle("/metrics", promhttp.Handler())
    	if err := http.ListenAndServe(addr, nil); err != nil {
        	panic(err)
    	}
	}()
}

func UpdateMemoryUsageGauge(memoryUsage prometheus.Gauge) {
	go func() {
		var memStats runtime.MemStats
		for {
			runtime.ReadMemStats(&memStats)
			memoryUsage.Set(float64(memStats.Alloc))
			time.Sleep(time.Second * 5)
		}
	}()
}
