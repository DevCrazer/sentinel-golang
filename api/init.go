package api

import (
	"fmt"
	"github.com/sentinel-group/sentinel-golang/core/config"
	"github.com/sentinel-group/sentinel-golang/core/log/metric"
	"github.com/sentinel-group/sentinel-golang/core/system"
	"github.com/sentinel-group/sentinel-golang/logging"
)

// InitDefault initializes Sentinel using the configuration from system
// environment and the default value.
func InitDefault() error {
	return initSentinel("", "")
}

// Init loads Sentinel general configuration from the given YAML file
// and initializes Sentinel. Note that the logging module will be initialized
// using the configuration from system environment or the default value.
func Init(configPath string) error {
	return initSentinel(configPath, "")
}

// InitWithLogDir initializes Sentinel logging module with the given directory.
// Then it loads Sentinel general configuration from the given YAML file
// and initializes Sentinel.
func InitWithLogDir(configPath string, logDir string) error {
	return initSentinel(configPath, logDir)
}

func initSentinel(configPath string, logDir string) (err error) {
	defer func() {
		if r := recover(); r != nil {
			var ok bool
			err, ok = r.(error)
			if !ok {
				err = fmt.Errorf("%v", r)
			}
		}
	}()

	// First we initialize the logging module.
	if logDir == "" {
		err = logging.InitializeLogConfigFromEnv()
	} else {
		err = logging.InitializeLogConfig(logDir, true)
	}
	if err != nil {
		return err
	}
	// Initialize the general configuration.
	err = config.InitConfigFromFile(configPath)
	if err != nil {
		return err
	}

	metric.InitTask()
	system.InitCollector()

	return err
}
