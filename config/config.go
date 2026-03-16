package config

import (
	"fmt"
	"os"
)

type Config struct {
	Telegram TelegramConfig
	Weather  WeatherConfig
	CalDAV   CalDAVConfig
	Digest   DigestConfig
}

type TelegramConfig struct {
	Token  string
	ChatID string
}

type WeatherConfig struct {
	APIKey string
	City   string
	Units  string
	Lang   string
}

type CalDAVConfig struct {
	URL      string
	Username string
	Password string
}

type DigestConfig struct {
	Cron string
}

func Load() (*Config, error) {
	cfg := &Config{
		Telegram: TelegramConfig{
			Token:  os.Getenv("TELEGRAM_BOT_TOKEN"),
			ChatID: os.Getenv("TELEGRAM_CHAT_ID"),
		},
		Weather: WeatherConfig{
			APIKey: os.Getenv("WEATHER_API_KEY"),
			City:   getEnvOrDefault("WEATHER_CITY", "Moscow"),
			Units:  getEnvOrDefault("WEATHER_UNITS", "metric"),
			Lang:   getEnvOrDefault("WEATHER_LANG", "ru"),
		},
		CalDAV: CalDAVConfig{
			URL:      os.Getenv("CALDAV_URL"),
			Username: os.Getenv("CALDAV_USERNAME"),
			Password: os.Getenv("CALDAV_PASSWORD"),
		},
		Digest: DigestConfig{
			Cron: getEnvOrDefault("DIGEST_CRON", "0 8 * * *"),
		},
	}

	if err := cfg.validate(); err != nil {
		return nil, err
	}

	return cfg, nil
}

func (c *Config) validate() error {
	if c.Telegram.Token == "" {
		return fmt.Errorf("TELEGRAM_BOT_TOKEN is required")
	}
	if c.Telegram.ChatID == "" {
		return fmt.Errorf("TELEGRAM_CHAT_ID is required")
	}
	if c.Weather.APIKey == "" {
		return fmt.Errorf("WEATHER_API_KEY is required")
	}
	return nil
}

func getEnvOrDefault(key, defaultVal string) string {
	if v := os.Getenv(key); v != "" {
		return v
	}
	return defaultVal
}
