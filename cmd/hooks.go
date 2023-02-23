/*
Copyright © 2023 Gabriele Puliti <gabriele.puliti+github@proton.me>
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"

	"imp/addons/gitlab"
)

var hookId int
var hookAction string

// hooksCmd represents the hook command
var hooksCmd = &cobra.Command{
	Use:   "hooks",
	Short: "Get list of gitlab project hooks",
	Long: `Get list of gitlab project hooks`,
    Run: func(cmd *cobra.Command, args []string) {
        switch hookAction {
        case "delete":
            if hookId != -1 {
                if gitlab.DeleteProjectHooksById(projectId, hookId) {
                    fmt.Println("Hook remove succesfully")
                } else {
                    fmt.Println("ERROR: Hook not remove")
                }
            } else {
                fmt.Println("Need a hook-id to delete a hook!")
            }
        default:
            for _, hook := range gitlab.GetProjectHooksById(projectId) {
                if hookId == -1 {
                    fmt.Println(hook)
                } else if hook.Id == hookId {
                    fmt.Println(hook)
                    break
                }
            }
        }
    },
}

func init() {
	projectCmd.AddCommand(hooksCmd)
    hooksCmd.PersistentFlags().IntVar(&hookId, "hook-id", -1, "gitlab hook id")
    hooksCmd.PersistentFlags().StringVar(&hookAction, "action", "", `Action to do with hook: 
    - delete
    // todo #10 : - update (with more args)
    `)
}
