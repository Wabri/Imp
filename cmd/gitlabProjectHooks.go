/*
Copyright © 2023 Gabriele Puliti <gabriele.puliti+github@proton.me>
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"

	"imp/addons/gitlab"
	"imp/utils/output"
)

var gitlabProjectHookId int
var gitlabProjectHookAction string

// hooksCmd represents the hook command
var gitlabProjectHooksCmd = &cobra.Command{
	Use:   "hooks",
	Short: "Get list of gitlab project hooks",
	Long: `Get list of gitlab project hooks`,
    Run: func(cmd *cobra.Command, args []string) {
        switch gitlabProjectHookAction {
        case "delete":
            if gitlabProjectHookId != -1 {
                if gitlab.DeleteProjectHooksById(gitlabProjectId, gitlabProjectHookId) {
                    fmt.Println("Hook remove succesfully")
                } else {
                    fmt.Println("ERROR: Hook not remove")
                }
            } else {
                fmt.Println("Need a hook-id to delete a hook!")
            }
        default:
            hooks := gitlab.GetProjectHooksById(gitlabProjectId)
            fmt.Println(output.AnyToString(hooks))
        }
    },
}

func init() {
	gitlabProjectCmd.AddCommand(gitlabProjectHooksCmd)
    gitlabProjectHooksCmd.PersistentFlags().IntVar(&gitlabProjectHookId, "hook-id", -1, "gitlab hook id")
    gitlabProjectHooksCmd.Flags().StringVar(&gitlabProjectHookAction, "action", "", `Action to do with hook: 
    - delete
    // todo #10 : - update (with more args)
    `)
}
