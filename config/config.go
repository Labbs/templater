package config

var AppConfig Config = Config{}

type Config struct {
	Version    string
	Debug      bool
	PrettyLogs bool

	ValuesFile    string
	TemplateFiles string
	OutputFiles   string
}
