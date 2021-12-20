package cmd

import (
	"fmt"
	"os"
	"strings"

	"github.com/spf13/cobra"
	"github.com/worldOneo/crudist"
	"github.com/worldOneo/crudist/core"
	"github.com/worldOneo/crudist/gen"
)

func init() {
	var configs []string
	var genCmd = &cobra.Command{
		Use:   "gen",
		Short: "Generates code from a config file",
		Long: `Generates code from a given config file.
The code is automatically written to files.
Previous files will be overwritten.`,
		RunE: func(cmd *cobra.Command, args []string) error {
			cli := fmt.Sprintf("crudist %s", strings.Join(os.Args[1:], " "))
			for _, file := range configs {
				conf, err := crudist.ReadConfig(file)
				if err != nil {
					return err
				}
				err = gen.Generate(core.Meta{CLIInput: cli}, conf)
				if err != nil {
					return err
				}
			}
			return nil
		},
	}
	rootCmd.AddCommand(genCmd)
	genCmd.Flags().StringArrayVarP(&configs, "config", "c", []string{"crudist.json"}, "The config file to generate code from")
}
