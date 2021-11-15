package main

import (
	"fmt"
	"os"
	"testing"

	"github.com/4rcode/gnomic/internal"
)

type TestRunOpt struct{}

func TestMain(t *testing.T) {
	is := internal.Assert(t)

	os.Args = []string{"test"}

	var arguments []interface{}

	fatal = func(args ...interface{}) {
		arguments = args
	}

	main()

	is(fmt.Sprint(arguments[0]) == "name cannot be empty", arguments[0])
}

func TestRun(t *testing.T) {
	is := internal.Assert(t)

	for tc, td := range map[string]struct {
		provided error
		expected interface{}
	}{
		"no flags": {
			Run(),
			"name cannot be empty",
		},
		"wrong flag": {
			Run("-z"),
			"flag provided but not defined: -z",
		},
		"help flag": {
			Run("-h"),
			nil,
		},
		"valid flags": {
			Run("-f", "optgen_", "-n", "TestRunOpt", "-p", "main", "-x", "_test"),
			nil,
		},
		"invalid package": {
			Run("-n", "TestRunOpt", "-p", "package"),
			"2:9: expected 'IDENT', found 'package' (and 1 more errors)",
		},
		"invalid file name": {
			Run("-n", "TestRunOpt", "-p", "main", "-F", string([]byte{0})),
			"open \x00: invalid argument",
		},
	} {
		t.Log(tc)

		is(fmt.Sprint(td.provided) == fmt.Sprint(td.expected), fmt.Sprint(td.provided))
	}

	templateText = "{{ .Invalid }}"

	expected := "template: :1:3: executing \"\" at <.Invalid>: " +
		"can't evaluate field Invalid in type main._flags"

	provided := fmt.Sprint(
		Run("-n", "TestRunOpt", "-p", "main"),
	)

	t.Log("invalid template")

	is(provided == expected, provided)
}
