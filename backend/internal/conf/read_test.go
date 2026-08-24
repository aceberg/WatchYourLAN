package conf

import (
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/aceberg/WatchYourLAN/internal/models"
	"github.com/spf13/viper"
)

func TestReadDefaultsConnectivityRetentionToTrimHistWhenAbsent(t *testing.T) {
	t.Cleanup(viper.Reset)
	viper.Reset()

	confPath := filepath.Join(t.TempDir(), "config_v2.yaml")
	if err := os.WriteFile(confPath, []byte("TRIM_HIST: 168\n"), 0o600); err != nil {
		t.Fatalf("os.WriteFile: %v", err)
	}

	config := read(confPath)

	if config.TrimHist != 168 {
		t.Fatalf("TrimHist = %d, want 168", config.TrimHist)
	}
	if config.ConnectivityRetention != 168 {
		t.Fatalf("ConnectivityRetention = %d, want TrimHist default 168", config.ConnectivityRetention)
	}
}

func TestConnectivityRetentionPersistsAndReadsBack(t *testing.T) {
	t.Cleanup(viper.Reset)
	viper.Reset()

	confPath := filepath.Join(t.TempDir(), "config_v2.yaml")
	if err := os.WriteFile(confPath, []byte("TRIM_HIST: 168\nCONNECTIVITY_RETENTION: 72\n"), 0o600); err != nil {
		t.Fatalf("os.WriteFile: %v", err)
	}

	config := read(confPath)
	if config.ConnectivityRetention != 72 {
		t.Fatalf("ConnectivityRetention = %d, want 72", config.ConnectivityRetention)
	}

	config.ConfPath = confPath
	config.ConnectivityRetention = 96
	if err := WriteErr(config); err != nil {
		t.Fatalf("WriteErr: %v", err)
	}

	written, err := os.ReadFile(confPath)
	if err != nil {
		t.Fatalf("os.ReadFile: %v", err)
	}
	if !strings.Contains(strings.ToLower(string(written)), "connectivity_retention: 96") {
		t.Fatalf("config file did not persist connectivity retention: %s", string(written))
	}

	viper.Reset()
	reread := read(confPath)
	if reread.ConnectivityRetention != 96 {
		t.Fatalf("reread ConnectivityRetention = %d, want 96", reread.ConnectivityRetention)
	}
}

func TestWriteErrPersistsConnectivityRetentionKey(t *testing.T) {
	t.Cleanup(viper.Reset)
	viper.Reset()

	confPath := filepath.Join(t.TempDir(), "config_v2.yaml")
	if err := os.WriteFile(confPath, []byte("{}\n"), 0o600); err != nil {
		t.Fatalf("os.WriteFile: %v", err)
	}

	config := models.Conf{
		ConfPath:              confPath,
		TrimHist:              48,
		ConnectivityRetention: 120,
	}

	if err := WriteErr(config); err != nil {
		t.Fatalf("WriteErr: %v", err)
	}

	written, err := os.ReadFile(confPath)
	if err != nil {
		t.Fatalf("os.ReadFile: %v", err)
	}
	if !strings.Contains(strings.ToLower(string(written)), "connectivity_retention: 120") {
		t.Fatalf("config file did not persist connectivity retention: %s", string(written))
	}
}
