package conf

import (
	"errors"
	"os"
	"path/filepath"
	"sync"
	"testing"

	"github.com/aceberg/WatchYourLAN/internal/models"
)

func TestUpdateAppConfigSerializesConcurrentUpdates(t *testing.T) {
	oldConfig := GetAppConfig()
	tempDir := t.TempDir()
	confPath := filepath.Join(tempDir, "config_v2.yaml")
	if err := os.WriteFile(confPath, []byte("{}\n"), 0o600); err != nil {
		t.Fatalf("os.WriteFile: %v", err)
	}
	SetAppConfigForTest(models.Conf{
		ConfPath: confPath,
		ArpStrs:  []string{},
	})
	t.Cleanup(func() {
		SetAppConfigForTest(oldConfig)
	})

	const updates = 20
	start := make(chan struct{})
	var wg sync.WaitGroup
	for i := 0; i < updates; i++ {
		wg.Add(1)
		go func(index int) {
			defer wg.Done()
			<-start
			_, err := UpdateAppConfig(func(config *models.Conf) error {
				config.ArpStrs = append(config.ArpStrs, string(rune('a'+index)))
				return nil
			})
			if err != nil {
				t.Errorf("UpdateAppConfig %d: %v", index, err)
			}
		}(i)
	}

	close(start)
	wg.Wait()

	config := GetAppConfig()
	if len(config.ArpStrs) != updates {
		t.Fatalf("ArpStrs len = %d, want %d: %+v", len(config.ArpStrs), updates, config.ArpStrs)
	}
}

func TestUpdateAppConfigDoesNotPersistRejectedMutation(t *testing.T) {
	oldConfig := GetAppConfig()
	tempDir := t.TempDir()
	confPath := filepath.Join(tempDir, "config_v2.yaml")
	if err := os.WriteFile(confPath, []byte("host: 127.0.0.1\n"), 0o600); err != nil {
		t.Fatalf("os.WriteFile: %v", err)
	}
	SetAppConfigForTest(models.Conf{
		ConfPath: confPath,
		Host:     "127.0.0.1",
	})
	t.Cleanup(func() {
		SetAppConfigForTest(oldConfig)
	})

	wantErr := errors.New("reject")
	_, err := UpdateAppConfig(func(config *models.Conf) error {
		config.Host = "0.0.0.0"
		return wantErr
	})
	if !errors.Is(err, wantErr) {
		t.Fatalf("UpdateAppConfig error = %v, want %v", err, wantErr)
	}

	if got := GetAppConfig().Host; got != "127.0.0.1" {
		t.Fatalf("Host = %q, want unchanged", got)
	}
	written, err := os.ReadFile(confPath)
	if err != nil {
		t.Fatalf("os.ReadFile: %v", err)
	}
	if string(written) != "host: 127.0.0.1\n" {
		t.Fatalf("config file changed after rejected mutation: %q", string(written))
	}
}
