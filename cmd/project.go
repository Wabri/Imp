/*
Copyright © 2023 Gabriele Puliti <gabriele.puliti+github@proton.me>

*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// projectCmd represents the project command
var projectCmd = &cobra.Command{
	Use:   "project",
	Short: "Get github project by id",
	Long: `Get github project by id`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("gitlab project")
	},
}

func init() {
	gitlabCmd.AddCommand(projectCmd)
}
