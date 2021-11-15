package main

import (
	"flag"
	"fmt"
	"os"
	"strings"
)

type _flags struct {
	Append, Omit, Replace, Skip bool

	Clear, ClearCondition,
	Default, DefaultCondition,
	fileName, filePrefix, fileSuffix,
	Mutator, mutatorPrefix, Name, Package string
}

func (f *_flags) parse(args ...string) error {
	pkg := os.Getenv("GOPACKAGE")

	var flagSet flag.FlagSet

	flagSet.BoolVar(&f.Append, "a", false, "append slice elements")
	flagSet.StringVar(&f.Clear, "c", "", "clear value")
	flagSet.StringVar(&f.ClearCondition, "C", "len(value) < 1", "clear condition")
	flagSet.StringVar(&f.Default, "d", "", "default value")
	flagSet.StringVar(&f.DefaultCondition, "D", "*target == nil", "default value condition")
	flagSet.StringVar(&f.fileName, "F", "", "file name, overrides -f and -x")
	flagSet.StringVar(&f.filePrefix, "f", "", "file name prefix")
	flagSet.StringVar(&f.fileSuffix, "x", "", "file name suffix")
	flagSet.StringVar(&f.Mutator, "M", "", "mutator name, overrides -m")
	flagSet.StringVar(&f.mutatorPrefix, "m", "Set", "mutator prefix")
	flagSet.StringVar(&f.Name, "n", "", "option name")
	flagSet.BoolVar(&f.Omit, "o", false, "omit array dereference on replace")
	flagSet.StringVar(&f.Package, "p", pkg, "package name")
	flagSet.BoolVar(&f.Replace, "r", false, "replace array/slice/map elements, overrides -a")
	flagSet.BoolVar(&f.Skip, "s", false, "skip mutator, overrides all flags")

	err := flagSet.Parse(args)

	if err != nil {
		return err
	}

	for _, flag := range []struct {
		name, value string
	}{
		{"name", f.Name},
		{"package", f.Package},
	} {
		if len(flag.value) < 1 {
			return fmt.Errorf("%s cannot be empty", flag.name)
		}
	}

	if len(f.Mutator) < 1 {
		f.Mutator = f.mutatorPrefix + f.Name
	}

	if len(f.fileName) < 1 {
		f.fileName = fmt.Sprintf(
			"%s%s%s.go", f.filePrefix,
			strings.ToLower(f.Name), f.fileSuffix)
	}

	return nil
}
