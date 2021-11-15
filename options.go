package gnomic

import (
	"sync"
)

// Options is a slice of Option instances.
type Options []Option

// Append returns the concatenation of the current Options with the provided
// options.
func (o Options) Append(options ...Option) Options {
	return append(o, options...)
}

// ApplyTo applies each Option to a target.
func (o Options) ApplyTo(target interface{}) {
	for _, option := range o {
		if option != nil {
			option.ApplyTo(target)
		}
	}
}

func (o Options) applyTo(target interface{}, wg *sync.WaitGroup) {
	defer wg.Done()

	o.ApplyTo(target)
}

// ApplyToAll applies each Option to all targets.
func (o Options) ApplyToAll(targets ...interface{}) {
	var wg sync.WaitGroup

	wg.Add(len(targets))

	for _, target := range targets {
		go o.applyTo(target, &wg)
	}

	wg.Wait()
}

// Parse returns Options with those values that are instances of Option.
func Parse(values ...interface{}) Options {
	length := len(values)

	if length < 1 {
		return nil
	}

	options := make(Options, 0, length)

	for _, value := range values {
		if option, ok := value.(Option); ok {
			options = options.Append(option)
		}
	}

	return options
}
