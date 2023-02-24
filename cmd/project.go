/*
Copyright © 2023 Gabriele Puliti <gabriele.puliti+github@proton.me>
*/
package cmd

import (
	"fmt"
	"imp/addons/gitlab"

	"github.com/spf13/cobra"
)

var projectId int
var projectAction string

// projectCmd represents the project command
var projectCmd = &cobra.Command{
	Use:   "project",
	Short: "Get github project by id",
	Long: `Get github project by id`,
	Run: func(cmd *cobra.Command, args []string) {
        switch projectAction {
        case "search-code":
            fmt.Println(gitlab.SearchOnProjectById(projectId, args[0]))
        default:
            fmt.Println(gitlab.GetProjectById(projectId)) 
        }
	},
}

func init() {
	gitlabCmd.AddCommand(projectCmd)
    projectCmd.PersistentFlags().IntVar(&projectId, "id", -1, "gitlab project id")
    projectCmd.PersistentFlags().StringVar(&projectAction, "action", "", `Action to do with hook: 
    - search-code
    `)
}
