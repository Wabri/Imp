/*
Copyright © 2023 Gabriele Puliti <gabriele.puliti+github@proton.me>
*/
package cmd

import (
	"encoding/json"
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
        if gitlabProjectId != -1 {
            switch gitlabProjectHookAction {
            case "delete":
                if gitlabProjectHookId != -1 {
                    if !gitlab.DeleteProjectHooksById(gitlabProjectId, gitlabProjectHookId) {
                        panic("ERROR: Hook not remove")
                    }
                } else {
                    panic("Need a hook-id to delete a hook!")
                }
            case "update":
                if gitlabProjectHookId != -1 {
                    var hook gitlab.Hook
                    json.Unmarshal([]byte(args[0]),&hook)
                    if !gitlab.PutProjectHooksById(gitlabProjectId, gitlabProjectHookId, hook) {
                        panic("ERROR: Hook not remove")
                    }
                } else {
                    panic("Need a hook-id to delete a hook!")
                }
            default:
                if gitlabProjectHookId != -1 {
                    hook := gitlab.GetProjectHookById(gitlabProjectId, gitlabProjectHookId)
                    fmt.Println(output.AnyToString(hook))
                } else {
                    hooks := gitlab.GetProjectHooksById(gitlabProjectId)
                    fmt.Println(output.AnyToString(hooks))
                }
        }
        }
    },
}

func init() {
	gitlabProjectCmd.AddCommand(gitlabProjectHooksCmd)
    gitlabProjectHooksCmd.PersistentFlags().IntVar(&gitlabProjectHookId, "hook-id", -1, "gitlab hook id")
    gitlabProjectHooksCmd.Flags().StringVar(&gitlabProjectHookAction, "action", "", `Action to do with hook: 
    - delete
    - update
    `)
}
