package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "learn-golang",
	Short: "Hello this is a CLI built using Cobra.",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Hello world!")
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
