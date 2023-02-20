/*
Copyright © 2023 Gabriele Puliti <gabriele.puliti+github@proton.me>
*/
package cmd

import (
	"fmt"
	"imp/addons/gitlab"

	"github.com/spf13/cobra"
)

var id int

// projectCmd represents the project command
var projectCmd = &cobra.Command{
	Use:   "project",
	Short: "Get github project by id",
	Long: `Get github project by id`,
	Run: func(cmd *cobra.Command, args []string) {
        fmt.Println(gitlab.GetProjectById(id)) 
	},
}

func init() {
	gitlabCmd.AddCommand(projectCmd)
    projectCmd.PersistentFlags().IntVar(&id, "id", -1, "gitlab project id")
}
