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

func tagValueEscaped(value string) string {
	value = strings.ReplaceAll(value, " ", "\\ ")
	value = strings.ReplaceAll(value, ",", "\\,")
	value = strings.ReplaceAll(value, "=", "\\=")

	return value
}

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

		// Escape special characters in tag values
		oneHist.Name = tagValueEscaped(oneHist.Name)
		if oneHist.Name == "" {
			oneHist.Name = "unknown"
		}

		networkName := tagValueEscaped(appConfig.InfluxNetworkName)
		deviceLocation := tagValueEscaped(appConfig.InfluxDeviceLocation)

		line := fmt.Sprintf("WatchYourLAN,IP=%s,iface=%s,name=%s,mac=%s,known=%d,network_name=%s,device_location=%s state=%d", oneHist.IP, oneHist.Iface, oneHist.Name, oneHist.Mac, oneHist.Known, networkName, deviceLocation, oneHist.Now)
		// slog.Debug("Writing to InfluxDB", "line", line)

		err = writeAPI.WriteRecord(context.Background(), line)
		check.IfError(err)
	} else {
		slog.Error("Can't connect to InfluxDB server")
		check.IfError(err)
	}

	client.Close()
}
