/*
Copyright © 2023 Gabriele Puliti <gabriele.puliti+github@proton.me>

*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// gitlabCmd represents the gitlab command
var gitlabCmd = &cobra.Command{
	Use:   "gitlab",
	Short: "Gitlab service management",
	Long: `Reach your gitlab service.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("gitlab called")
	},
}

func init() {
	rootCmd.AddCommand(gitlabCmd)
}
