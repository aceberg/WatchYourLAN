package migrate

import (
	"fmt"
	"log"

	"github.com/aceberg/WatchYourLAN/internal/check"
	"github.com/aceberg/WatchYourLAN/internal/conf"
	"github.com/aceberg/WatchYourLAN/internal/notify"
)

// ToYAML - check if config file is YAML and migrate from ENV-file if not
func ToYAML(path string) string {

	if check.IsYaml(path) {
		if check.IsEmpty(path) {

			oldPath := path[:len(path)-5] // path w/o .yaml

			if !check.IsEmpty(oldPath) {
				migrateConfig(oldPath, path)
			}
		}
	} else {
		newPath := path + ".yaml"
		check.Path(newPath)

		if check.IsEmpty(newPath) && !check.IsEmpty(path) {

			migrateConfig(path, newPath)
		}

		path = newPath
	}

	return path
}

func migrateConfig(oldPath, newPath string) {

	config := getOldConfig(oldPath)
	conf.Write(newPath, config)

	msg := fmt.Sprintf("Config file migrated from ENV to YAML format. \nOld config file %s is no longer used and can be deleted. \nNew config file is %s, please update your settings.", oldPath, newPath)

	log.Println("=================================== ")
	log.Println("WARNING:", msg)
	log.Println("=================================== ")
	notify.Shoutrrr("WatchYourLAN: "+msg, config.ShoutURL)
}
