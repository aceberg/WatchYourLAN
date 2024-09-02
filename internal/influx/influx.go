package influx

import (
	"context"
	"crypto/tls"
	"fmt"
	"log/slog"
	"strings"

	"github.com/influxdata/influxdb-client-go/v2"

	"github.com/aceberg/WatchYourLAN/internal/check"
	"github.com/aceberg/WatchYourLAN/internal/models"
)

// Add - write data to InfluxDB2
func Add(appConfig models.Conf, oneHist models.Host) {
	var ctx context.Context

	client := influxdb2.NewClientWithOptions(appConfig.InfluxAddr, appConfig.InfluxToken,
		influxdb2.DefaultOptions().
			SetUseGZip(true).
			SetTLSConfig(&tls.Config{
				InsecureSkipVerify: appConfig.InfluxSkipTLS,
			}))

	ctx = context.Background()
	ping, err := client.Ping(ctx)
	if ping {
		writeAPI := client.WriteAPIBlocking(appConfig.InfluxOrg, appConfig.InfluxBucket)

		// Escape special characters in strings
		oneHist.Name = strings.ReplaceAll(oneHist.Name, " ", "\\ ")
		oneHist.Name = strings.ReplaceAll(oneHist.Name, ",", "\\,")
		oneHist.Name = strings.ReplaceAll(oneHist.Name, "=", "\\=")

		line := fmt.Sprintf("WatchYourLAN,IP=%s,iface=%s,name=%s,mac=%s,known=%d state=%d", oneHist.IP, oneHist.Iface, oneHist.Name, oneHist.Mac, oneHist.Known, oneHist.Now)
		// slog.Debug("Writing to InfluxDB", "line", line)

		err = writeAPI.WriteRecord(context.Background(), line)
		check.IfError(err)
	} else {
		slog.Error("Can't connect to InfluxDB server")
		check.IfError(err)
	}

	client.Close()
}
