package conf

import (
	"github.com/caarlos0/env/v6"
	"time"
)

type App struct {
	PrometheusBind  string        `env:"PROMETHEUS_BIND" envDefault:":2112"`
	BotToken        string        `env:"BOT_TOKEN" envDefault:""`
	ChannelID       string        `env:"CHANNEL_ID" envDefault:""`
	RefreshDuration time.Duration `env:"REFRESH_DURATION" envDefault:"1m"`
}

func ParseEnv() (*App, error) {
	cfg := App{}
	err := env.Parse(&cfg)
	if err != nil {
		return nil, err
	}
	return &cfg, nil
}
