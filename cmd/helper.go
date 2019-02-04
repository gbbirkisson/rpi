package cmd

import (
	"fmt"
	"os"
	"sort"
	"strings"

	"github.com/olekukonko/tablewriter"
	"github.com/spf13/viper"
)

// ExitOnError prints the provided message and error, and then exits the program with status code 1
func ExitOnError(msg string, err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, msg+": %v\n", err)
		os.Exit(1)
	}
}

// PrintConfig prints to the Stdout all viper configuration
func PrintConfig() {
	m := viper.AllSettings()
	keys := make([]string, 0, len(m))
	for k := range m {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	table := tablewriter.NewWriter(os.Stdout)
	table.SetBorder(false)
	table.SetCenterSeparator("")
	table.SetColumnSeparator("")
	table.SetRowSeparator("")
	table.SetAlignment(tablewriter.ALIGN_LEFT)

	for _, k := range keys {
		v := []string{k, fmt.Sprintf("%v", m[k])}
		table.Append(v)
	}
	table.Render()
}

// PrintVersion prints name, version and revision to Stdout
func PrintVersion(name, version, revision string) {
	v := strings.Trim(fmt.Sprintf("%s %s", version, revision), " ")
	fmt.Printf("%s version %s \n", name, v)
}
