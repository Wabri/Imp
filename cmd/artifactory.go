/*
Copyright © 2023 Gabriele Puliti <gabriele.puliti+github@proton.me>
*/
package cmd

import (
	"fmt"
	"imp/addons/artifactory"
	"imp/utils/output"

	"github.com/spf13/cobra"
)

// gitlabCmd represents the gitlab command
var artifactoryCmd = &cobra.Command{
	Use:   "artifactory",
	Short: "Artifactory service management",
	Long: `Reach your artifactory service.`,
	Run: artifactoryRun,
}

func init() {
	rootCmd.AddCommand(artifactoryCmd)
}

func artifactoryRun(cmd *cobra.Command, args []string) {
    fmt.Println(output.AnyToString(artifactory.RequestHandler))
}

