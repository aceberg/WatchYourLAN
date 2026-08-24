package conf

import (
	"sync"

	"github.com/aceberg/WatchYourLAN/internal/check"
	"github.com/aceberg/WatchYourLAN/internal/models"
)

// AppConfig - app config
var AppConfig models.Conf
var appConfigMu sync.RWMutex

// Start - initial config
func Start(dirPath, nodePath string) {

	confPath := dirPath + "/config_v2.yaml"
	check.Path(confPath)

	config := read(confPath)

	config.DirPath = dirPath
	config.ConfPath = confPath
	config.DBPath = dirPath + "/scan.db"
	if nodePath != "" {
		config.NodePath = nodePath
	}

	setAppConfig(config)
}

// GetAppConfig returns a synchronized immutable snapshot of the active config.
func GetAppConfig() models.Conf {
	appConfigMu.RLock()
	defer appConfigMu.RUnlock()

	return cloneConfig(AppConfig)
}

// UpdateAppConfig serializes a read-modify-persist-write update to avoid lost changes.
func UpdateAppConfig(mutator func(*models.Conf) error) (models.Conf, error) {
	appConfigMu.Lock()
	defer appConfigMu.Unlock()

	nextConfig := cloneConfig(AppConfig)
	if err := mutator(&nextConfig); err != nil {
		return cloneConfig(AppConfig), err
	}
	if err := writeErrNoLock(nextConfig); err != nil {
		return cloneConfig(AppConfig), err
	}

	AppConfig = cloneConfig(nextConfig)
	return cloneConfig(AppConfig), nil
}

// SetVersion updates the runtime version without persisting it to the config file.
func SetVersion(version string) {
	appConfigMu.Lock()
	defer appConfigMu.Unlock()

	AppConfig.Version = version
}

// SetAppConfigForTest replaces the process config. Runtime code should use Start or UpdateAppConfig.
func SetAppConfigForTest(config models.Conf) {
	setAppConfig(config)
}

func setAppConfig(config models.Conf) {
	appConfigMu.Lock()
	defer appConfigMu.Unlock()

	AppConfig = cloneConfig(config)
}

func cloneConfig(config models.Conf) models.Conf {
	if config.ArpStrs != nil {
		config.ArpStrs = append([]string(nil), config.ArpStrs...)
	}

	return config
}
