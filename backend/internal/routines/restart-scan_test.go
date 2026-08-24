package routines

import (
	"sync"
	"testing"

	"github.com/aceberg/WatchYourLAN/internal/conf"
)

func TestScanRestartConcurrentDoesNotPanic(t *testing.T) {
	scanRestartMu.Lock()
	oldQuitScan := quitScan
	oldStartScanFunc := startScanFunc
	oldLogLevel := conf.AppConfig.LogLevel
	quitScan = make(chan bool)
	startScanFunc = func(chan bool) {}
	conf.AppConfig.LogLevel = "info"
	scanRestartMu.Unlock()

	t.Cleanup(func() {
		scanRestartMu.Lock()
		quitScan = oldQuitScan
		startScanFunc = oldStartScanFunc
		conf.AppConfig.LogLevel = oldLogLevel
		scanRestartMu.Unlock()
	})

	var wg sync.WaitGroup
	panics := make(chan any, 32)

	for i := 0; i < 32; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			defer func() {
				if r := recover(); r != nil {
					panics <- r
				}
			}()

			ScanRestart()
		}()
	}

	wg.Wait()
	close(panics)

	for panicValue := range panics {
		t.Fatalf("ScanRestart panicked: %v", panicValue)
	}
}
