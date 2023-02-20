/*
Copyright © 2023 Gabriele Puliti <gabriele.puliti+github@proton.me>
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"

	"imp/addons/gitlab"
)

var hook_id int
var action string

// hooksCmd represents the hook command
var hooksCmd = &cobra.Command{
	Use:   "hooks",
	Short: "Get list of gitlab project hooks",
	Long: `Get list of gitlab project hooks`,
    Run: func(cmd *cobra.Command, args []string) {
        switch action {
        case "delete":
            if hook_id != -1 {
                if gitlab.DeleteProjectHooksById(id, hook_id) {
                    fmt.Println("Hook remove succesfully")
                } else {
                    fmt.Println("ERROR: Hook not remove")
                }
            } else {
                fmt.Println("Need a hook-id to delete a hook!")
            }
        default:
            for _, hook := range gitlab.GetProjectHooksById(id) {
                if hook_id == -1 {
                    fmt.Println(hook)
                } else if hook.Id == hook_id {
                    fmt.Println(hook)
                    break
                }
            }
        }
    },
}

func init() {
	projectCmd.AddCommand(hooksCmd)
    hooksCmd.PersistentFlags().IntVar(&hook_id, "hook-id", -1, "gitlab hook id")
    hooksCmd.PersistentFlags().StringVar(&action, "action", "", `Action to do with hook: 
    - delete
    // todo #10 : - update (with more args)
    `)
}
