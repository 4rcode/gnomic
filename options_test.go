package gnomic_test

import (
	"reflect"
	"testing"

	"github.com/4rcode/gnomic"
	"github.com/4rcode/gnomic/internal"
)

type A string

//go:generate go run ./cmd/optgen -f optgen_ -x _test -n A

type B int

//go:generate go run ./cmd/optgen -f optgen_ -x _test -n B

func TestOptionsAppend(t *testing.T) {
	is := internal.Assert(t)

	type opts = gnomic.Options

	for tc, td := range map[string]struct {
		provided, expected opts
	}{
		"nil options": {
			nil,
			opts{nil, B(123)},
		},
		"empty options": {
			opts{},
			opts{nil, B(123)},
		},
		"mixed options": {
			opts{nil, A("abc")},
			opts{nil, A("abc"), nil, B(123)},
		},
	} {
		t.Log(tc)

		provided := td.provided.Append().Append(nil, B(123))

		is(reflect.DeepEqual(provided, td.expected), td.provided)
	}
}

func TestOptionsApplyToAndApplyToAll(t *testing.T) {
	is := internal.Assert(t)

	type opts = gnomic.Options

	for tc, td := range map[string]struct {
		opts
		a, x A
		b, y B
	}{
		"nil options": {
			nil,
			"xyz", "xyz", 1, 2,
		},
		"empty options": {
			opts{},
			"xyz", "xyz", 1, 2,
		},
		"mixed options": {
			opts{nil, A("abc"), B(123)},
			"abc", "xyz", 123, 2,
		},
	} {
		t.Log(tc)

		var a, x A = "xyz", "xyz"
		var b, y B = 1, 2

		td.opts.ApplyTo(nil)
		td.opts.ApplyTo(A("xyz"))
		td.opts.ApplyTo(&a)
		td.opts.ApplyTo(x)
		td.opts.ApplyToAll()
		td.opts.ApplyToAll(nil, A("xyz"), B(321), &b, y)

		is(a == td.a, a)
		is(b == td.b, b)
		is(x == td.x, x)
		is(y == td.y, y)
	}
}

func TestParse(t *testing.T) {
	is := internal.Assert(t)

	type opts = gnomic.Options

	for tc, td := range map[string]struct {
		provided, expected opts
	}{
		"nil options": {
			nil,
			opts{B(123)},
		},
		"empty options": {
			opts{},
			opts{B(123)},
		},
		"mixed options": {
			opts{nil, A("abc")},
			opts{nil, A("abc"), B(123)},
		},
	} {
		t.Log(tc)

		provided := td.provided.Append(
			gnomic.Parse()...,
		).Append(
			gnomic.Parse(nil, 123, B(123))...,
		)

		is(reflect.DeepEqual(provided, td.expected), provided)
	}
}
