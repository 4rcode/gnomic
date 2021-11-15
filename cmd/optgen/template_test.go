package main

import (
	"testing"
	"time"

	"github.com/4rcode/gnomic"
	"github.com/4rcode/gnomic/internal"
)

type Opt01 time.Duration

//go:generate go run . -f optgen_ -x _test -n Opt01

type Opt02 int

//go:generate go run . -f optgen_ -x _test -n Opt02 -s

type Opt03 [1]int

//go:generate go run . -f optgen_ -x _test -n Opt03 -F optgen_opt03_file_test.go

type Opt04 [1]int

//go:generate go run . -f optgen_ -x _test -n Opt04 -o -c Opt04{1} -C "value[0] < 0"

type Opt05 []int

//go:generate go run . -f optgen_ -x _test -n Opt05 -m Apply

type Opt06 []int

//go:generate go run . -f optgen_ -x _test -n Opt06 -a -m abc -M SetValue

type Opt07 []int

//go:generate go run . -f optgen_ -x _test -n Opt07 -a -c nil

type Opt08 []int

//go:generate go run . -f optgen_ -x _test -n Opt08 -a -r -d "make(Opt08, len(value))" -D "len(*target) < len(value)"

type Opt09 []int

//go:generate go run . -f optgen_ -x _test -n Opt09 -r -c nil -d "make(Opt09, len(value))" -D "len(*target) < len(value)"

type Opt10 map[int]int

//go:generate go run . -f optgen_ -x _test -n Opt10 -p main

type Opt11 map[int]int

//go:generate go run . -f optgen_ -x _test -n Opt11 -a -r -d Opt11{}

type Opt12 map[int]int

//go:generate go run . -f optgen_ -x _test -n Opt12 -r -c nil -d Opt12{}

func TestOptions(t *testing.T) {
	is := internal.Assert(t)

	var cfg struct {
		Opt01
		Opt02
		Opt03
		Opt04
		Opt05
		Opt06
		Opt07
		Opt08
		Opt09
		Opt10
		Opt11
		Opt12
	}

	gnomic.Options{
		Opt01(time.Nanosecond),
		Opt02(2),
		Opt03{123},
		Opt04{123},
		Opt05{1, 2, 3}, Opt05{}, Opt05{123},
		Opt06{1, 2, 3}, Opt06{}, Opt06{123}, Opt06{456},
		Opt07{1, 2, 3}, Opt07{}, Opt07{123}, Opt07{456},
		Opt08{1, 2, 3}, Opt08{}, Opt08{123}, Opt08{456},
		Opt09{1, 2, 3}, Opt09{}, Opt09{123}, Opt09{456},
		Opt10{1: 11, 2: 22, 3: 33}, Opt10{}, Opt10{123: 123}, Opt10{456: 456},
		Opt11{1: 11, 2: 22, 3: 33}, Opt11{}, Opt11{123: 123}, Opt11{456: 456},
		Opt12{1: 11, 2: 22, 3: 33}, Opt12{}, Opt12{123: 123}, Opt12{456: 456},
	}.ApplyTo(&cfg)

	is(cfg.Opt01 == 1, cfg.Opt01)
	is(cfg.Opt02 == 20, cfg.Opt02)
	is(cfg.Opt03[0] == 123, cfg.Opt03)
	is(cfg.Opt04[0] == 123, cfg.Opt04)

	if is(len(cfg.Opt05) == 1, cfg.Opt05) {
		is(cfg.Opt05[0] == 123, cfg.Opt05)
	}

	if is(len(cfg.Opt06) == 5, cfg.Opt06) {
		is(cfg.Opt06[0] == 1, cfg.Opt06)
		is(cfg.Opt06[1] == 2, cfg.Opt06)
		is(cfg.Opt06[2] == 3, cfg.Opt06)
		is(cfg.Opt06[3] == 123, cfg.Opt06)
		is(cfg.Opt06[4] == 456, cfg.Opt06)
	}

	if is(len(cfg.Opt07) == 2, cfg.Opt07) {
		is(cfg.Opt07[0] == 123, cfg.Opt07)
		is(cfg.Opt07[1] == 456, cfg.Opt07)
	}

	if is(len(cfg.Opt08) == 3, cfg.Opt08) {
		is(cfg.Opt08[0] == 456, cfg.Opt08)
		is(cfg.Opt08[1] == 2, cfg.Opt08)
		is(cfg.Opt08[2] == 3, cfg.Opt08)
	}

	if is(len(cfg.Opt09) == 1, cfg.Opt09) {
		is(cfg.Opt09[0] == 456, cfg.Opt09)
	}

	if is(len(cfg.Opt10) == 1, cfg.Opt10) {
		is(cfg.Opt10[456] == 456, cfg.Opt10)
	}

	if is(len(cfg.Opt11) == 5, cfg.Opt11) {
		is(cfg.Opt11[1] == 11, cfg.Opt11)
		is(cfg.Opt11[2] == 22, cfg.Opt11)
		is(cfg.Opt11[3] == 33, cfg.Opt11)
		is(cfg.Opt11[123] == 123, cfg.Opt11)
		is(cfg.Opt11[456] == 456, cfg.Opt11)
	}

	if is(len(cfg.Opt12) == 2, cfg.Opt12) {
		is(cfg.Opt12[123] == 123, cfg.Opt12)
		is(cfg.Opt12[456] == 456, cfg.Opt12)
	}
}
