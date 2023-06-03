package main

import (
	"fmt"
	"io"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	filtercmd "src/cmd/smgr/cmd/filter"
)

func NewRootCommand(output io.Writer) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "smgr",
		Short: "Manage Semantic Versioning compliant versions.",
		Long:  `Manage Semantic Versioning compliant versions and integrate with popular repository and registry platform to facilitate the task.`,
		PersistentPreRun: func(cmd *cobra.Command, args []string) {

			viper.SetConfigName("ccs")
			viper.AddConfigPath(".")
			viper.SetEnvPrefix("ccs")
			viper.AutomaticEnv()
			if err := viper.ReadInConfig(); err != nil {
				cmd.Println("Warning reading config file:", err)
			}
		},
	}

	filterCmd := filtercmd.NewFilterCommand()
	cmd.AddCommand(filterCmd)

	return cmd
}

func main() {
	rootCmd := NewRootCommand(nil)
	// Execute the root command
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}