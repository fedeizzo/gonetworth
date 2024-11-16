package config

import (
	"fmt"
	"net/url"
	"os"

	"github.com/rs/zerolog"
	"gopkg.in/yaml.v3"
)

type LogFormat string

const (
	FormatJson LogFormat = "json"
	FormatText LogFormat = "text"
)

type Config struct {
	LogLevel  zerolog.Level `yaml:"log_level"`
	LogFormat LogFormat     `yaml:"log_format"`

	Postgres struct {
		Host     string `yaml:"host"`
		Port     string `yaml:"port"`
		User     string `yaml:"user"`
		Password string `yaml:"password"`
		DB       string `yaml:"db"`
		Args     string `yaml:"args"`
	} `yaml:"postgres"`
}

func Parse() Config {
	cfg := Config{}

	configFile, err := os.ReadFile("./gonetworth.yaml")
	if err != nil {
		panic(err)
	}
	if err := yaml.Unmarshal(configFile, &cfg); err != nil {
		panic(err)
	}

	return cfg
}

func (c Config) DatabaseUrl() (*url.URL, error) {
	credentials := ""
	if c.Postgres.User != "" {
		if c.Postgres.Password != "" {
			credentials = fmt.Sprintf("%s:%s@", c.Postgres.User, c.Postgres.Password)
		} else {
			credentials = fmt.Sprintf("%s@", c.Postgres.User)
		}
	}

	u, err := url.Parse(fmt.Sprintf("postgres://%s%s:%s/%s?%s", credentials, c.Postgres.Host, c.Postgres.Port, c.Postgres.DB, c.Postgres.Args))
	if err != nil {
		return nil, err
	}

	return u, nil
}

func (c Config) RedactedDatabaseUrl() (*url.URL, error) {
	credentials := ""
	if c.Postgres.User != "" {
		if c.Postgres.Password != "" {
			credentials = fmt.Sprintf("%s:%s@", c.Postgres.User, "***")
		} else {
			credentials = fmt.Sprintf("%s@", c.Postgres.User)
		}
	}

	u, err := url.Parse(fmt.Sprintf("postgres://%s%s:%s/%s?%s", credentials, c.Postgres.Host, c.Postgres.Port, c.Postgres.DB, c.Postgres.Args))
	if err != nil {
		return nil, err
	}

	return u, nil
}
