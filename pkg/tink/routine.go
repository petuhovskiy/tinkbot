package tink

import (
	"github.com/petuhovskiy/go-template/pkg/conf"
	"github.com/petuhovskiy/telegram"
	log "github.com/sirupsen/logrus"
	"time"
)

type state struct {
	bot      *telegram.Bot
	cfg      *conf.App
	lastRate float64
}

func (s *state) tick() error {
	rate, err := FetchExchangeRate(s.cfg.FromCurrency, s.cfg.ToCurrency)
	if err != nil {
		return err
	}

	msg, err := s.bot.SendMessage(&telegram.SendMessageRequest{
		ChatID: s.cfg.ChannelID,
		Text:   rate.FormattedString,
	})
	if err != nil {
		return err
	}

	isUpdated := rate.Rate != s.lastRate
	s.lastRate = rate.Rate

	if isUpdated {
		_, err = s.bot.PinChatMessage(&telegram.PinChatMessageRequest{
			ChatID:    s.cfg.ChannelID,
			MessageID: msg.MessageID,
		})
		if err != nil {
			return err
		}
	}

	return nil
}

func RefreshRoutine(bot *telegram.Bot, cfg *conf.App) {
	ticker := time.NewTicker(cfg.RefreshDuration)

	state := state{
		bot:      bot,
		cfg:      cfg,
		lastRate: 0,
	}

	for range ticker.C {
		err := state.tick()
		if err != nil {
			log.WithError(err).Error("RefreshRoutine tick error")
		}
	}
}
