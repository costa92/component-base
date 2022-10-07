package main

import (
	"flag"
	"fmt"
	"os"
	"strings"

	"github.com/gosuri/uitable"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
)

const configFlagName = "server"

var (
	cfgFile string
	age     int
)

// nolint
func init() {
	pflag.StringVarP(&cfgFile, "server", "c", cfgFile, "Read configuration from specified `FILE`, support JSON, TOML, YAML, HCL, or Java properties formats.")
	pflag.IntVarP(&age, "age", "a", age, "age")
}

func wordSepNormalizeFunc(_ *pflag.FlagSet, name string) pflag.NormalizedName {
	if strings.Contains(name, "_") {
		return pflag.NormalizedName(strings.ReplaceAll(name, "_", "-"))
	}
	return pflag.NormalizedName(name)
}

func main() {
	basename := "costa_long"
	pflag.CommandLine.SetNormalizeFunc(wordSepNormalizeFunc)
	pflag.CommandLine.AddGoFlagSet(flag.CommandLine)
	pflag.Lookup("age").NoOptDefVal = "11"
	pflag.Lookup(configFlagName).NoOptDefVal = "dev.yaml"
	pflag.Parse()
	fmt.Println(cfgFile)
	fmt.Println(age)
	fmt.Println(&age)

	viper.SetConfigFile(cfgFile)
	viper.SetEnvPrefix(strings.ReplaceAll(strings.ToUpper(basename), "-", "_"))
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_", "-", "_"))

	if err := viper.ReadInConfig(); err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "Error: failed to read configuration file(%s): %v\n", cfgFile, err)
		os.Exit(1)
	}

	keys := viper.AllKeys()
	if len(keys) > 0 {
		table := uitable.New()
		table.RightAlign(0)
		for _, k := range keys {
			table.AddRow(fmt.Sprintf("%s:", k), viper.Get(k))
		}
		fmt.Printf("%v", table)
	}
}
