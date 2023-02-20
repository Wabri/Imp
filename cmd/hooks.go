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

// hooksCmd represents the hook command
var hooksCmd = &cobra.Command{
	Use:   "hooks",
	Short: "Get list of gitlab project hooks",
	Long: `Get list of gitlab project hooks`,
    Run: func(cmd *cobra.Command, args []string) {
        for _, hook := range gitlab.GetProjectHooksById(id) {
            if hook_id == -1 {
                fmt.Println(hook)
            } else if hook.Id == hook_id {
                fmt.Println(hook)
                break
            }
        }
    },
}

func init() {
	projectCmd.AddCommand(hooksCmd)
    hooksCmd.PersistentFlags().IntVar(&hook_id, "hook-id", -1, "gitlab hook id")
}
