package main

import (
	"fmt"
	"os"

	"github.com/BurntSushi/toml"
	"github.com/vbauerster/zaplizer"
	"go.uber.org/zap"
)

type config struct {
	LogMode  string
	LogLevel string
}

func main() {
	cfg, err := loadConfig("config.toml")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	logger, err := zaplizer.InitLogger(zaplizer.ToLogMode(cfg.LogMode), zaplizer.ToLogLevel(cfg.LogLevel))
	if err != nil {
		logger = zap.L()
	}
	logger.Info("test")
	logger.Warn("test")
	logger.Debug("test")
}

func loadConfig(path string) (*config, error) {
	cfg := new(config)
	if _, err := toml.DecodeFile(path, cfg); err != nil {
		return cfg, fmt.Errorf("cannot load config: %v", err)
	}
	return cfg, nil
}
