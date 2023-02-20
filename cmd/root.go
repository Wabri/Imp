/*
Copyright © 2023 Gabriele Puliti <gabriele.puliti@proton.me>
*/
package cmd

import (
	"fmt"
	"os"

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
        fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
}

