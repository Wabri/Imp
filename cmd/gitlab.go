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

// gitlabCmd represents the gitlab command
var gitlabCmd = &cobra.Command{
	Use:   "gitlab",
	Short: "Gitlab service management",
	Aliases: []string{"g", "gl"},
	Run: gitlabRun,
}

func init() {
	rootCmd.AddCommand(gitlabCmd)
}

func gitlabRun(cmd *cobra.Command, args []string) {
    fmt.Println(output.AnyToString(gitlab.RequestHandler))
}
