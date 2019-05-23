package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use: "ethquery",
}

var verboseFlag bool = false

func Execute() {
	rootCmd.PersistentFlags().BoolVarP(&verboseFlag, "verbose", "v", false, "Verbose output")

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
