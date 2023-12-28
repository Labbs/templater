package config

var AppConfig Config = Config{}

type Config struct {
	ValuesFile    string
	TemplateFiles string
	OutputFiles   string
}
