package metrics

import (
    "github.com/prometheus/client_golang/prometheus"
)

var (
    MemoryUsage = prometheus.NewGauge(
        prometheus.GaugeOpts{
            Name: "memory_usage_bytes",
            Help: "Current memory usage in bytes",
        },
    )
    
    fileReadProgress = prometheus.NewGauge(
        prometheus.GaugeOpts{
            Name: "file_read_progress_percent",
            Help: "Percentage of file read",
        },
    )
    
    linesProcessed = prometheus.NewCounter(
        prometheus.CounterOpts{
            Name: "lines_processed_total",
            Help: "Total number of lines processed",
        },
    )
    
    lineProcessingDuration = prometheus.NewHistogram(
        prometheus.HistogramOpts{
            Name:    "line_processing_duration_seconds",
            Help:    "Histogram of time spent processing lines",
            Buckets: prometheus.DefBuckets,
        },
    )
)
