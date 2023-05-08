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

var groupId int
var groupName string
var groupAction string

var gitlabGroupCmd = &cobra.Command{
	Use:   "group",
	Short: "Manage gitlab group",
	Long: "Manage gitlab group",
	Example: `
- Default: Get all the groups
    imp gitlab group
- action: projects -> List of projects for a given group 
    imp gitlab project --id 1234 --action projects
	`,
	Run: func(cmd *cobra.Command, args []string) {
        switch groupAction {
        case "projects":
	    projects := gitlab.GetProjectsByGroup(groupId)
            fmt.Println(output.AnyToString(projects))
        default:
	    if groupId > 0 {
		group := gitlab.GetGroupById(groupId)
		fmt.Println(output.AnyToString(group))
	    } else if groupName != "" {
		result := gitlab.SearchGroup(groupName)
		fmt.Println(output.AnyToString(result[0]))
	    } else {
		groups := gitlab.GetGroups()
		fmt.Println(output.AnyToString(groups))
	    }
        }
	},
}

func init() {
	gitlabCmd.AddCommand(gitlabGroupCmd)
	gitlabGroupCmd.PersistentFlags().StringVar(&groupName, "name", "", "gitlab group name")
	gitlabGroupCmd.PersistentFlags().IntVar(&groupId, "id", -1, "gitlab group id")
	gitlabGroupCmd.Flags().StringVar(&groupAction, "action", "", "Possible actions: projects")
}
