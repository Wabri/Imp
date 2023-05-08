/*
Copyright © 2023 Gabriele Puliti <gabriele.puliti+github@proton.me>
*/
package cmd

import (
	"fmt"
	"imp/addons/gitlab"
	"imp/utils/output"

	"github.com/spf13/cobra"
)

var gitlabProjectId int
var gitlabProjectAction string

// projectCmd represents the project command
var gitlabProjectCmd = &cobra.Command{
	Use:   "project",
	Short: "Get github project by id",
	Long: `Get github project by id`,
	Example: `
- Default: Get basic info for a project using id
    imp gitlab project --id 1234
- action: search-code -> Search for code
    imp gitlab project --id 1234 --action search-code test
- action: settings -> Change settings
    imp gitlab project --id 1234 --action settings visibility internal
	`,
	Run: func(cmd *cobra.Command, args []string) {
        switch gitlabProjectAction {
        case "settings":
            result := gitlab.ChangeSettingById(gitlabProjectId, args[0], args[1])
            fmt.Println(output.AnyToString(result))
        case "search-code":
            project := gitlab.SearchOnProjectById(gitlabProjectId, args[0])
            fmt.Println(output.AnyToString(project))
        default:
            project := gitlab.GetProjectById(gitlabProjectId)
            fmt.Println(output.AnyToString(project))
        }
	},
}

func init() {
	gitlabCmd.AddCommand(gitlabProjectCmd)
	gitlabProjectCmd.PersistentFlags().IntVar(&gitlabProjectId, "id", -1, "gitlab project id")
	gitlabProjectCmd.Flags().StringVar(&gitlabProjectAction, "action", "", "Possible actions: search-code, settings")
}
