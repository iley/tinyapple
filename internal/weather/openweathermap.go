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

// OpenWeatherMapProvider is an implementation of a weather Provider using OpenWeatherMap API.
type OpenWeatherMapProvider struct {
	mu           sync.Mutex
	apiKey       string
	locationName string
	current      string
}

// NewOpenWeatherMapProvider creates an initializes an instance of OpenWeatherMapProvider.
func NewOpenWeatherMapProvider(ctx context.Context, apiKey, locationName string) *OpenWeatherMapProvider {
	p := &OpenWeatherMapProvider{
		apiKey:       apiKey,
		locationName: locationName,
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
	updateInterval := time.Hour // TODO: Consider making this a configurable option.
	for {
		err := p.update(ctx)
		if err != nil {
			log.Infof("Could not fetch weather information from OpenWeatherMap: %s", err)
		}
		select {
		case <-ctx.Done():
			return
		case <-time.After(updateInterval):
			continue
		}
	}
}

func (p *OpenWeatherMapProvider) update(ctx context.Context) error {
	cwd, err := openweathermap.NewCurrent("C", "en", p.apiKey)
	if err != nil {
		return err
	}
	if err = cwd.CurrentByName(p.locationName); err != nil {
		return err
	}
	newCurrent := getWeatherString(cwd)
	p.mu.Lock()
	p.current = newCurrent
	p.mu.Unlock()
	return nil
}

func getWeatherString(cwd *openweathermap.CurrentWeatherData) string {
	temp := cwd.Main.Temp
	prefix := ""
	if temp > 0 {
		prefix = "+"
	}
	return fmt.Sprintf("%s%.f°C", prefix, temp)
}
