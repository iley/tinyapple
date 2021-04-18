package weather

// Provider provides information about weather.
type Provider interface {
	Current() string
}
