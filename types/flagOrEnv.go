package types

import (
	"fmt"
	"os"
	"strings"

	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
)

type FlagOrEnv struct {
	name      string
	value     string
	envPrefix string
	defVal    string

	noOptFlag bool
}

func (f *FlagOrEnv) String() string {
	val := f.value
	if val == "" || val == f.defVal {
		envVal := os.Getenv(fmt.Sprintf("%s%s", f.envPrefix, f.name))
		if envVal != "" {
			val = envVal
		}
	}
	return val
}

func (f *FlagOrEnv) Set(value string) error {
	if f.noOptFlag {
		f.value = "true"
		return nil
	}
	delimiter := ","
	if f.value == "" {
		delimiter = ""
	}
	f.value = fmt.Sprintf("%s%s%s", f.value, delimiter, value)
	return nil
}

func (f *FlagOrEnv) Type() string {
	return "FlagOrEnv"
}

func (f *FlagOrEnv) AddToCobraCommand(cmd *cobra.Command, name, short, defVal, usage, prefix string, isNoOpt bool) {
	if cmd == nil {
		return
	}

	f.name = normalize(name)
	f.envPrefix = normalize(prefix)
	f.noOptFlag = isNoOpt
	f.value = defVal
	f.defVal = defVal

	flag := &pflag.Flag{
		Name:      name,
		Shorthand: short,
		Usage:     usage,
		Value:     f,
		DefValue:  defVal,
	}

	if isNoOpt {
		flag.NoOptDefVal = defVal
	}

	cmd.Flags().AddFlag(flag)
}

func normalize(in string) string {
	out := strings.Replace(in, "-", "_", -1)

	return strings.ToUpper(out)
}
