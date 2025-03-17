package prometheus

import (
	"strconv"

	"github.com/aceberg/WatchYourLAN/internal/models"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

var up *prometheus.GaugeVec

func NewMetrics() {
	up = promauto.NewGaugeVec(prometheus.GaugeOpts{
		Namespace: "watch_your_lan",
		Name:      "up",
		Help:      "Whether the host is up (1 for yes, 0 for no)",
	}, []string{"ip", "iface", "name", "mac", "known"})
}

// Add a Prometheus metric
func Add(oneHist models.Host) {
	if up == nil {
		return
	}

	if oneHist.Name == "" {
		oneHist.Name = "unknown"
	}

	up.With(prometheus.Labels{
		"ip":    oneHist.IP,
		"iface": oneHist.Iface,
		"name":  oneHist.Name,
		"mac":   oneHist.Mac,
		"known": strconv.Itoa(oneHist.Known),
	}).Set(float64(oneHist.Now))
}
