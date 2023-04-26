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

var artifactoryCmd = &cobra.Command{
	Use:   "artifactory",
	Short: "Artifactory service management",
	Aliases: []string{"arti", "a"},
	Example: `
	- Default: Get the list of all artifacts on specified repository
	    imp artifactory --repository test-local
	    imp artifactory --repository test-local --artifact test/local/1.2.3/
	- action: repository-list -> Get the list of all repositories
	    imp artifactory --action repository-list
	- action: artifact-list -> Get the list of all artifacts on specified repository
	    imp artifactory --action artifact-list test-local/test/local/1.2.3
	    imp artifactory --repository test-local --action artifact-list 
	- action: not-used-since -> List all the artifact of a repository not used by a specified date
	    imp artifactory --repository test-local --action not-used-since 2023-03-14
	- action: delete-artifact -> Remove the artifact on a specified repository
	    imp artifactory --repository test-local --artifact test/local/1.2.3/artifact.json --action delete-artifact
	- action: get-info -> Get all the infos about specific repository or artifact
	    imp artifactory --repository test-local --action get-info
	    imp artifactory --repository test-local --action get-info test/local/1.2.3
	`,
	Run: artifactoryRun,
}

func init() {
	rootCmd.AddCommand(artifactoryCmd)
	artifactoryCmd.Flags().StringVar(&artifact, "artifact", "", `Set artifact`)
	artifactoryCmd.Flags().StringVar(&repository, "repository", "", `Set repository`)
	artifactoryCmd.Flags().StringVar(&artifactoryAction, "action", "", "Possible actions: repository-list, artifact-list, not-used-since, delete-artifact, get-info")
}

func artifactoryRun(cmd *cobra.Command, args []string) {
	switch artifactoryAction {
	case "repository-list":
		repositories := artifactory.GetRepositories()
		fmt.Println(output.AnyToString(repositories))
		return
	case "artifact-list":
		if !valid(repository) { repository = args[0] }
		artifacts := artifactory.GetArtifacts(repository)
		fmt.Println(output.AnyToString(artifacts))
		return
	case "not-used-since":
		if !valid(repository) { return }
		notUsedSince := args[0]
		artifacts := artifactory.GetArtifactNotUsedSinceForRepository(repository, notUsedSince)
		fmt.Println(output.AnyToString(artifacts))
		return
	case "delete-artifact":
		if !valid(artifact) { return }
		if !valid(repository) { return }
		result := artifactory.DeleteArtifact(repository, artifact)
		if (result) { 
			fmt.Println("INFO: Deleted") 
		} else {
			fmt.Println("ERR: Unable to delete") 
		}
		return
	case "get-info":
		if !valid(repository) { return }
		path := ""
		if len(args) != 0 { path = args[0]}
		items := artifactory.GetItemInfos(repository, path)
		fmt.Println(output.AnyToString(items))
		return
	default:
		if !valid(repository) { 
			fmt.Println("ERR: At least a repository must be passed") 
			return 
		}
		path := repository + "/" + artifact
		artifact := artifactory.GetArtifacts(path)
		fmt.Println(output.AnyToString(artifact))
		return
	}
}

func valid(element string) bool {
	if len(element) == 0 {
		fmt.Println("[Err]: All the necessary parameters must be pass")
		return false
	}
	return true
}
