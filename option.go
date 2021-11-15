package gnomic

// Option represents a configured value.
type Option interface {
	// ApplyTo is responsible for applying a configured value to a target.
	ApplyTo(target interface{})
}
