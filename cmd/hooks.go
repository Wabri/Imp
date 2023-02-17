/*
Copyright © 2023 Gabriele Puliti <gabriele.puliti+github@proton.me>

*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// hooksCmd represents the hook command
var hooksCmd = &cobra.Command{
	Use:   "hooks",
	Short: "Get list of gitlab project hooks",
	Long: `Get list of gitlab project hooks`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("hooks called")
	},
}

func init() {
	projectCmd.AddCommand(hooksCmd)
}
