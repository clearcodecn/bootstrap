package config

import (
	"github.com/ghodss/yaml"
	"log"
	"os"
)

var (
	_config *Config
)

type Config struct {
	Driver string `json:"driver"`
	Dsn    string `json:"dsn"`

	CachePath           string `json:"cachePath"`
	Cache               bool   `json:"cache"`
	CacheDuration       int    `json:"cacheDuration"`
	StaticCacheDuration int    `json:"staticCacheDuration"`
	ListenAddress       string `json:"listenAddress"`

	DevId        string `json:"devId"`
	HeaderScript string `json:"headerScript"`
	FooterScript string `json:"footerScript"`
	TemplatePath string `json:"templatePath"`
	StaticPath   string `json:"staticPath"`

	Hosts []*HostConfig `json:"hosts"`
	Debug bool          `json:"debug"`
}

type HostConfig struct {
	Id           string `json:"id"`
	Host         string `json:"host"`
	AdsTxt       string `json:"adsTxt"`
	Theme        string `json:"theme"`
	Title        string `json:"title"`
	Description  string `json:"description"`
	Keywords     string `json:"keywords"`
	CdnUrl       string `json:"cdnUrl"`
	AdsList      Ad     `json:"adsList"`
	HeaderScript string `json:"headerScript"`
	FooterScript string `json:"footerScript"`
}

type Ad struct {
	Link     string `json:"link"`
	Type     string `json:"type"`
	Text     string `json:"text"`
	Image    string `json:"image"`
	Direct   bool   `json:"direct"`
	Position string `json:"position"`
}

func Get() *Config {
	return _config
}

func GetHost(cid string) *HostConfig {
	c := Get()
	for _, h := range c.Hosts {
		if h.Id == cid {
			return h
		}
	}
	panic("invalid cid: " + cid)
}

func Parse(configfile string) {
	data, err := os.ReadFile(configfile)
	if err != nil {
		log.Fatal("读取配置失败" + err.Error())
	}
	var cfg Config
	err = yaml.Unmarshal(data, &cfg)
	if err != nil {
		log.Fatal("读取配置失败" + err.Error())
	}
	_config = &cfg
}
