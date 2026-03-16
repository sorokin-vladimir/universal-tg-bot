package digest

import (
	"fmt"
	"strings"
	"time"

	"github.com/sorokin-vladimir/universal-tg-bot/internal/calendar"
	"github.com/sorokin-vladimir/universal-tg-bot/internal/weather"
)

type Service struct {
	weather  *weather.Client
	calendar *calendar.Client
}

func New(w *weather.Client, c *calendar.Client) *Service {
	return &Service{weather: w, calendar: c}
}

func (s *Service) Build() (string, error) {
	var sb strings.Builder

	sb.WriteString(fmt.Sprintf("*Доброе утро! %s*\n\n", time.Now().Format("02.01.2006, Monday")))

	forecast, err := s.weather.Today()
	if err != nil {
		sb.WriteString("Погода: не удалось получить данные\n")
	} else {
		sb.WriteString("*Погода сегодня*\n")
		sb.WriteString(fmt.Sprintf("%s\n", forecast.Description))
		sb.WriteString(fmt.Sprintf("Температура: %.0f°C (ощущается %.0f°C)\n", forecast.TempCurrent, forecast.FeelsLike))
		sb.WriteString(fmt.Sprintf("Мин/Макс: %.0f°C / %.0f°C\n", forecast.TempMin, forecast.TempMax))
		sb.WriteString(fmt.Sprintf("Влажность: %d%%\n", forecast.Humidity))
		sb.WriteString(fmt.Sprintf("Ветер: %.1f м/с\n", forecast.WindSpeed))
	}

	// Calendar section will be populated in step 2
	events, err := s.calendar.EventsRange(7)
	if err == nil && len(events) > 0 {
		sb.WriteString("\n*События на ближайшие 7 дней*\n")
		for _, e := range events {
			sb.WriteString(fmt.Sprintf("• %s — %s\n", e.Start, e.Title))
		}
	}

	return sb.String(), nil
}
