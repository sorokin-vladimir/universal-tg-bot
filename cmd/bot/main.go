package main

import (
	"log"
	"os"
	"os/signal"
	"strconv"
	"syscall"

	"github.com/joho/godotenv"

	"github.com/sorokin-vladimir/universal-tg-bot/config"
	"github.com/sorokin-vladimir/universal-tg-bot/internal/bot"
	"github.com/sorokin-vladimir/universal-tg-bot/internal/calendar"
	"github.com/sorokin-vladimir/universal-tg-bot/internal/digest"
	"github.com/sorokin-vladimir/universal-tg-bot/internal/scheduler"
	"github.com/sorokin-vladimir/universal-tg-bot/internal/weather"
)

func main() {
	_ = godotenv.Load()

	cfg, err := config.Load()
	if err != nil {
		log.Fatalf("config: %v", err)
	}

	chatID, err := strconv.ParseInt(cfg.Telegram.ChatID, 10, 64)
	if err != nil {
		log.Fatalf("invalid TELEGRAM_CHAT_ID: %v", err)
	}

	tg, err := bot.New(cfg.Telegram.Token)
	if err != nil {
		log.Fatalf("telegram bot: %v", err)
	}

	weatherClient := weather.NewClient(
		cfg.Weather.APIKey,
		cfg.Weather.City,
		cfg.Weather.Units,
		cfg.Weather.Lang,
	)

	calClient := calendar.NewClient(
		cfg.CalDAV.URL,
		cfg.CalDAV.Username,
		cfg.CalDAV.Password,
	)

	digestSvc := digest.New(weatherClient, calClient)

	sched := scheduler.New()
	if err := sched.Add(cfg.Digest.Cron, func() {
		text, err := digestSvc.Build()
		if err != nil {
			log.Printf("digest build error: %v", err)
			return
		}
		if err := tg.SendMessage(chatID, text); err != nil {
			log.Printf("send message error: %v", err)
		}
	}); err != nil {
		log.Fatalf("schedule digest: %v", err)
	}

	sched.Start()
	log.Printf("bot started, digest scheduled: %s", cfg.Digest.Cron)

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	sched.Stop()
	log.Println("bot stopped")
}
