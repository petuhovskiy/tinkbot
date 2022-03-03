package main

import (
	"encoding/json"
	"github.com/davecgh/go-spew/spew"
	"github.com/petuhovskiy/go-template/pkg/tink"
	"github.com/petuhovskiy/telegram"
	"github.com/petuhovskiy/telegram/updates"
	"net/http"

	"github.com/petuhovskiy/go-template/pkg/conf"

	"github.com/prometheus/client_golang/prometheus/promhttp"
	log "github.com/sirupsen/logrus"
)

func main() {
	log.SetFormatter(&log.JSONFormatter{})
	log.SetReportCaller(true)
	log.SetLevel(log.DebugLevel)

	cfg, err := conf.ParseEnv()
	if err != nil {
		log.WithError(err).Fatal("failed to parse config from env")
	}

	go func() {
		mux := http.NewServeMux()
		mux.Handle("/metrics", promhttp.Handler())
		err := http.ListenAndServe(cfg.PrometheusBind, mux)
		if err != nil && err != http.ErrServerClosed {
			log.WithError(err).Fatal("prometheus server error")
		}
	}()

	bot := telegram.NewBotWithOpts(cfg.BotToken, &telegram.Opts{
		Middleware: func(handler telegram.RequestHandler) telegram.RequestHandler {
			return func(methodName string, req interface{}) (message json.RawMessage, err error) {
				res, err := handler(methodName, req)
				if err != nil {
					log.Println("Telegram response error: ", err)
				}

				return res, err
			}
		},
	})

	ch, err := updates.StartPolling(bot, telegram.GetUpdatesRequest{
		Offset:  0,
		Limit:   50,
		Timeout: 10,
	})
	if err != nil {
		log.Fatal(err)
	}

	go tink.RefreshRoutine(bot, cfg)

	for upd := range ch {
		spew.Dump(upd)
	}
}
