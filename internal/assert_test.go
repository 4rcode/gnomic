package internal_test

import (
	"fmt"
	"strings"

	"github.com/4rcode/gnomic/internal"
)

type _context string

func (_context) Helper() {
	fmt.Println("helper")
}

func (_context) Errorf(format string, args ...interface{}) {
	fmt.Println(
		strings.ReplaceAll(
			fmt.Sprintf(format, args...), "\n", ""))
}

func Example() {
	var context _context

	is := internal.Assert(context)

	is(true, 123)
	is(false, "abc")

	// Output:
	// helper
	// unexpected "abc"
}
