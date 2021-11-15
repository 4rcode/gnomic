# gnomic - configuration library

[![
	Build
](https://github.com/4rcode/gnomic/actions/workflows/build.yml/badge.svg)
](https://github.com/4rcode/gnomic/actions/workflows/build.yml)
[![
	Coverage
](https://codecov.io/gh/4rcode/gnomic/branch/main/graph/badge.svg)
](https://codecov.io/gh/4rcode/gnomic/branch/main)
[![
	Reference
](https://pkg.go.dev/badge/github.com/4rcode/gnomic.svg)
](https://pkg.go.dev/github.com/4rcode/gnomic)

## Installation

```sh
go install github.com/4rcode/gnomic/cmd/optgen@latest
go get github.com/4rcode/gnomic
```

## [Quick Start](quickstart_test.go)

```go
import (
	"fmt"
	"time"

	"github.com/4rcode/gnomic"
)

type Message string

//go:generate optgen -n Message

func Send(opts ...gnomic.Option) {
	cfg := struct {
		Message
	}{
		Message: "nothing to send",
	}

	gnomic.
		Options(opts).
		ApplyTo(&cfg)

	fmt.Println(cfg.Message)
}

func Example() {
	Send(Message("hi!"))
}
```
