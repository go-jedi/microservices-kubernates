package config

import (
	"flag"
	"log"
	"os"

	"gopkg.in/yaml.v3"
)

type LoggerConfig struct {
	IsJSON     bool   `yaml:"is_json"`
	AddSource  bool   `yaml:"add_source"`
	Level      string `yaml:"level"`
	SetFile    bool   `yaml:"set_file"`
	FileName   string `yaml:"file_name"`
	MaxSize    int    `yaml:"max_size"`
	MaxBackups int    `yaml:"max_backups"`
	MaxAge     int    `yaml:"max_age"`
}

type ClientConfig struct {
	TimeoutReq int `yaml:"timeout_req"`
	Posts      struct {
		URL string `yaml:"url"`
	} `yaml:"posts"`
}

type CorsConfig struct {
	AllowOrigins        []string `yaml:"allow_origins"`
	AllowMethods        []string `yaml:"allow_methods"`
	AllowHeaders        []string `yaml:"allow_headers"`
	ExposeHeaders       []string `yaml:"expose_headers"`
	MaxAge              int      `yaml:"max_age"`
	AllowCredentials    bool     `yaml:"allow_credentials"`
	AllowPrivateNetwork bool     `yaml:"allow_private_network"`
}

type HTTPServerConfig struct {
	Host              string     `yaml:"host"`
	Port              int        `yaml:"port"`
	EnablePrefork     bool       `yaml:"enable_prefork"`
	EnablePrintRoutes bool       `yaml:"enable_print_routes"`
	Cors              CorsConfig `yaml:"cors"`
}

type Config struct {
	Logger     LoggerConfig     `yaml:"logger"`
	Client     ClientConfig     `yaml:"client"`
	HTTPServer HTTPServerConfig `yaml:"httpserver"`
}

// LoadConfig load config file.
func LoadConfig() string {
	var cf string

	flag.StringVar(&cf, "config", "config.yaml", "config file path")
	flag.Parse()

	return cf
}

// ParseConfig parse config file.
func ParseConfig(configFile string) (config Config, err error) {
	f, err := os.Open(configFile)
	if err != nil {
		return config, err
	}
	defer func(f *os.File) {
		if err := f.Close(); err != nil {
			log.Printf("error closing the file: %v", err)
		}
	}(f)

	err = yaml.NewDecoder(f).Decode(&config)

	return config, err
}

// GetConfig get config.
func GetConfig() (config Config, err error) {
	cf := LoadConfig()

	return ParseConfig(cf)
}
