/*
Copyright © 2023 Gabriele Puliti <gabriele.puliti@proton.me>
*/
package cmd

import (
	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "imp",
	Short: "A cli written by an IMPostor on IMPulse that is unlikely to be IMPortant to anyone",
	Long: `Imp is a cli that help you to reach and manage services without using graphical interface or complex curl command.`,
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
        panic(err)
	}
}

func init() {
}

