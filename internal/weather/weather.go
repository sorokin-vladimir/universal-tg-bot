package weather

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
)

const apiURL = "https://api.openweathermap.org/data/2.5/weather"

type Client struct {
	apiKey string
	city   string
	units  string
	lang   string
	http   *http.Client
}

type Forecast struct {
	Description string
	TempCurrent float64
	TempMin     float64
	TempMax     float64
	FeelsLike   float64
	Humidity    int
	WindSpeed   float64
}

type apiResponse struct {
	Weather []struct {
		Description string `json:"description"`
	} `json:"weather"`
	Main struct {
		Temp      float64 `json:"temp"`
		FeelsLike float64 `json:"feels_like"`
		TempMin   float64 `json:"temp_min"`
		TempMax   float64 `json:"temp_max"`
		Humidity  int     `json:"humidity"`
	} `json:"main"`
	Wind struct {
		Speed float64 `json:"speed"`
	} `json:"wind"`
}

func NewClient(apiKey, city, units, lang string) *Client {
	return &Client{
		apiKey: apiKey,
		city:   city,
		units:  units,
		lang:   lang,
		http:   &http.Client{},
	}
}

func (c *Client) Today() (*Forecast, error) {
	params := url.Values{}
	params.Set("q", c.city)
	params.Set("appid", c.apiKey)
	params.Set("units", c.units)
	params.Set("lang", c.lang)

	resp, err := c.http.Get(apiURL + "?" + params.Encode())
	if err != nil {
		return nil, fmt.Errorf("weather api request: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("weather api returned status %d", resp.StatusCode)
	}

	var data apiResponse
	if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
		return nil, fmt.Errorf("decode weather response: %w", err)
	}

	forecast := &Forecast{
		TempCurrent: data.Main.Temp,
		TempMin:     data.Main.TempMin,
		TempMax:     data.Main.TempMax,
		FeelsLike:   data.Main.FeelsLike,
		Humidity:    data.Main.Humidity,
		WindSpeed:   data.Wind.Speed,
	}
	if len(data.Weather) > 0 {
		forecast.Description = data.Weather[0].Description
	}

	return forecast, nil
}
