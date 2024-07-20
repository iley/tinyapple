package weather

import (
	"context"
	"fmt"
	"sync"
	"time"

	"github.com/briandowns/openweathermap"
	log "github.com/sirupsen/logrus"
)

var _ Provider = (*OpenWeatherMapProvider)(nil)

const UmbrellaSymbol = "☂"

// OpenWeatherMapProvider is an implementation of a weather Provider using OpenWeatherMap API.
type OpenWeatherMapProvider struct {
	mu             sync.Mutex
	apiKey         string
	locationLat		 float64
	locationLon		 float64
	current        string
	updateInterval time.Duration
}

// NewOpenWeatherMapProvider creates an initializes an instance of OpenWeatherMapProvider.
func NewOpenWeatherMapProvider(ctx context.Context, apiKey string, lat, lon float64) *OpenWeatherMapProvider {
	p := &OpenWeatherMapProvider{
		apiKey:       apiKey,
		locationLat: lat,
		locationLon: lon,
		// Initially update frequently because newtork might not be available immediately on boot.
		updateInterval: 30 * time.Second,
	}
	go p.updateLoop(ctx)
	return p
}

// Current returns information about current weather as a string.
func (p *OpenWeatherMapProvider) Current() string {
	p.mu.Lock()
	defer p.mu.Unlock()
	if p.current == "" {
		return "..."
	}
	return p.current
}

func (p *OpenWeatherMapProvider) updateLoop(ctx context.Context) {
	for {
		err := p.update(ctx)
		if err == nil {
			// Slow down requests after first successful update.
			p.updateInterval = 15 * time.Minute
		}
		if err != nil {
			log.Warnf("Could not fetch weather information from OpenWeatherMap: %s", err)
		}
		select {
		case <-ctx.Done():
			return
		case <-time.After(p.updateInterval):
			continue
		}
	}
}

func (p *OpenWeatherMapProvider) update(_ context.Context) error {
	log.Debugf("Fetching weather information from OpenWeatherMap lat=%f lon=%f", p.locationLat, p.locationLon)

	weatherData, err := openweathermap.NewOneCall("C", "en", p.apiKey, []string{
		openweathermap.ExcludeMinutely,
		openweathermap.ExcludeHourly,
		openweathermap.ExcludeAlerts,
	})

	if err != nil {
		return err
	}

	if err = weatherData.OneCallByCoordinates(&openweathermap.Coordinates{Latitude: p.locationLat, Longitude: p.locationLon}); err != nil {
		return err
	}
	
	log.Debugf("Done fetching weather information: %v", weatherData)

	newCurrent := getWeatherString(weatherData)
	p.mu.Lock()
	p.current = newCurrent
	p.mu.Unlock()

	return nil
}

func getWeatherString(weather *openweathermap.OneCallData) string {
	temp := weather.Current.Temp
	prefix := ""
	if temp > 0 {
		prefix = "+"
	}

	// If it's going to rain at any point today, show an umbrella symbol.
	rainIndicator := ""
	if weather.Daily == nil || len(weather.Daily) == 0 {
		log.Warn("No daily weather data available")
		rainIndicator = "?"
	} else if weather.Daily[0].Rain > 0 {
		rainIndicator = UmbrellaSymbol
	}

	return fmt.Sprintf("%s%.f°C%s", prefix, temp, rainIndicator)
}
