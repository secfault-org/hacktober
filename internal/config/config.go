package config

import (
	"github.com/charmbracelet/log"
	"github.com/secfault-org/hacktober/internal/model/container"
	"github.com/spf13/viper"
)

type SSHConfig struct {
	Host      string
	Port      string
	PublicURL string
	KeyPath   string
}

type ContainerConfig struct {
	ContainerPort    container.Port
	TimeoutInMinutes int
}

type ChallengeConfig struct {
	BaseDir string
}

type Config struct {
	SSH       SSHConfig
	Container ContainerConfig
	Challenge ChallengeConfig
}

func NewConfig(c *viper.Viper) *Config {
	var cfg Config

	err := c.Unmarshal(&cfg)
	if err != nil {
		log.Fatal("Could not unmarshal config", "error", err)
	}

	return &cfg
}
