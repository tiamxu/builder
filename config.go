package main

import (
	"github.com/sirupsen/logrus"
	"github.com/tiamxu/kit/log"

	"github.com/koding/multiconfig"
	httpkit "github.com/tiamxu/kit/http"
)

const configPath = "config/config.yaml"

// yaml文件内容映射到结构体
type Config struct {
	ENV      string                  `yaml:"env"`
	LogLevel string                  `yaml:"log_level"`
	HttpSrv  httpkit.GinServerConfig `yaml:"http_srv"`
}

// set log level
func (c *Config) Initial() (err error) {
	defer func() {
		if err == nil {
			log.Printf("config initialed, env: %s", cfg.ENV)
		}
	}()

	if level, err := logrus.ParseLevel(c.LogLevel); err != nil {
		return err
	} else {
		log.DefaultLogger().SetLevel(level)
	}

	return nil
}

// 读取配置文件
func loadConfig() {
	cfg = new(Config)
	multiconfig.MustLoadWithPath(configPath, cfg)
}
