package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "crudist",
	Short: "Codegeneration for Go CRUD API",
	Long: `Crudist is a powerfull API tool to automate the generation
of boilberplate code which is often seen.`,
	TraverseChildren: true,
}


// Execute runs rootCmd
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}