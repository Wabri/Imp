/*
Copyright © 2023 Gabriele Puliti <gabriele.puliti+github@proton.me>
*/
package cmd

import (
	"fmt"
	"imp/addons/gitlab"

	"github.com/spf13/cobra"
)

var gitlabProjectId int
var gitlabProjectAction string

// projectCmd represents the project command
var gitlabProjectCmd = &cobra.Command{
	Use:   "project",
	Short: "Get github project by id",
	Long: `Get github project by id`,
	Run: func(cmd *cobra.Command, args []string) {
        switch gitlabProjectAction {
        case "search-code":
            fmt.Println(gitlab.SearchOnProjectById(gitlabProjectId, args[0]))
        default:
            fmt.Println(gitlab.GetProjectById(gitlabProjectId)) 
        }
	},
}

func init() {
	gitlabCmd.AddCommand(gitlabProjectCmd)
    gitlabProjectCmd.PersistentFlags().IntVar(&gitlabProjectId, "id", -1, "gitlab project id")
    gitlabProjectCmd.Flags().StringVar(&gitlabProjectAction, "action", "", `Action to do with hook: 
    - search-code
    `)
}
