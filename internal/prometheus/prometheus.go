package prometheus

import (
	"net/http"
	"strconv"

	"github.com/aceberg/WatchYourLAN/internal/models"
	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func Handler(appConfig *models.Conf) func(c *gin.Context) {
	h := promhttp.Handler()
	return func(c *gin.Context) {
		if !appConfig.PrometheusEnable {
			c.AbortWithStatus(http.StatusNotFound)
			return
		}
		h.ServeHTTP(c.Writer, c.Request)
	}
}

var up *prometheus.GaugeVec

func NewMetrics() {
	up = promauto.NewGaugeVec(prometheus.GaugeOpts{
		Namespace: "watch_your_lan",
		Name:      "up",
		Help:      "Whether the host is up (1 for yes, 0 for no)",
	}, []string{"ip", "iface", "name", "mac", "known"})
}

func RemoveMetrics() {
	up = nil
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
