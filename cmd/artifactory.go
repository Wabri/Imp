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

var artifactoryAction string
var artifact string
var repository string

// gitlabCmd represents the gitlab command
var artifactoryCmd = &cobra.Command{
	Use:   "artifactory",
	Short: "Artifactory service management",
	Long: `Reach your artifactory service.`,
	Run: artifactoryRun,
}

func init() {
	rootCmd.AddCommand(artifactoryCmd)
	artifactoryCmd.Flags().StringVar(&artifact, "", "", `Set artifact`)
	artifactoryCmd.PersistentFlags().StringVar(&repository, "", "", `Set repository`)
	artifactoryCmd.Flags().StringVar(&artifactoryAction, "artifactory", "", `Action to do with hook: 
	- not-used-since -> List all the artifact of a repository not used by a specified date
	    e.g.: imp artifactory --repository test-local --action not-used-since 2023-03-14
	- delete-artifact -> Remove the artifact on a specified repository
	    e.g.: imp artifactory --repository test-local --artifact test/local/1.2.3/artifact.json --action delete-artifact
	`)
}

func artifactoryRun(cmd *cobra.Command, args []string) {
	switch artifactoryAction {
	case "not-used-since":
		validate(repository)
		notUsedSince := args[0]
		repositories := artifactory.GetArtifactNotUsedSinceForRepository(repository, notUsedSince)
		fmt.Println(output.AnyToString(repositories))
		return
	case "delete-artifact":
		validate(repository)
		validate(artifact)
		result := artifactory.DeleteArtifact(repository, artifact)
		if (result) { 
			fmt.Println("INFO: Deleted") 
		} else {
			fmt.Println("ERR: Unable to delete") 
		}
		return
	default:
		validate(artifact)
		artifactInfo := artifactory.GetArtifactByName(artifact)
		fmt.Println(output.AnyToString(artifactInfo))
		return
	}
}

func validate(element string)  {
	if len(element) == 0 {
		fmt.Println("[Err]: All the necessary parameters must be pass")
		return
	}
}
