package gnomic_test

import (
	"fmt"
	"time"

	"github.com/4rcode/gnomic"
)

//
// Generate gnomic.Option implementations with optgen.
//
// e.g.
//
//  type Message string
//  type Timeout time.Duration
//
//  //go:generate optgen -n Message
//  //go:generate optgen -n Timeout
//

type Message string
type Timeout time.Duration

//go:generate go run ./cmd/optgen -f optgen_ -x _test -n Message
//go:generate go run ./cmd/optgen -f optgen_ -x _test -n Timeout

//
// Accept instances of gnomic.Option in your functions.
//
// Alternatively you can accept
//   ...interface{}
// then use Parse().
//

func Send(opts ...gnomic.Option) {

	//
	// Declare your settings.
	//

	var msg Message = "hi!"
	var tmt Timeout = 123

	//
	// Group your options using a struct, if you prefer.
	//

	cfg := struct {
		Message
		Timeout
	}{
		Message: "WoW! It took less than %s!",
		Timeout: 123,
	}

	//
	// Apply the options to the desired targets.
	//

	gnomic.
		Options(opts).
		ApplyToAll(&msg, &tmt, &cfg)

	//
	// Start using the values.
	//

	fmt.Println(msg)
	fmt.Println(tmt)

	message := string(cfg.Message)
	duration := time.Duration(cfg.Timeout)

	fmt.Printf(message, duration)

}

func Example() {

	Send(Timeout(100 * time.Millisecond))

	// Output:
	// hi!
	// 100000000
	// WoW! It took less than 100ms!

}
