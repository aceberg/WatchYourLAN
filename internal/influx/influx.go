package influx

import (
	"context"
	"crypto/tls"
	"fmt"
	// "log"
	"strings"

	"github.com/influxdata/influxdb-client-go/v2"

	"github.com/aceberg/WatchYourLAN/internal/check"
	"github.com/aceberg/WatchYourLAN/internal/models"
)

// Add - write data to InfluxDB2
func Add(appConfig models.Conf, oneHist models.Host) {

	client := influxdb2.NewClientWithOptions(appConfig.InfluxAddr, appConfig.InfluxToken,
		influxdb2.DefaultOptions().
			SetUseGZip(true).
			SetTLSConfig(&tls.Config{
				InsecureSkipVerify: appConfig.InfluxSkipTLS,
			}))

	writeAPI := client.WriteAPIBlocking(appConfig.InfluxOrg, appConfig.InfluxBucket)

	// Escape special characters in strings
	oneHist.Name = strings.ReplaceAll(oneHist.Name, " ", "\\ ")
	oneHist.Name = strings.ReplaceAll(oneHist.Name, ",", "\\,")
	oneHist.Name = strings.ReplaceAll(oneHist.Name, "=", "\\=")

	line := fmt.Sprintf("WatchYourLAN,IP=%s,iface=%s,name=%s,mac=%s,known=%d state=%d", oneHist.IP, oneHist.Iface, oneHist.Name, oneHist.Mac, oneHist.Known, oneHist.Now)
	// log.Println("LINE:", line)

	err := writeAPI.WriteRecord(context.Background(), line)
	check.IfError(err)

	client.Close()
}
