/*
Copyright © 2023 Gabriele Puliti <gabriele.puliti+github@proton.me>
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"

	"imp/addons/artifactory"
	"imp/utils/output"
)

var isRepositoriesList bool
var notUsedSince string
var deleteArtifact string

var artifactoryRepositoryCmd = &cobra.Command{
	Use:   "repository",
	Short: "Get list of repositories",
	Long: `Get list of repositories`,
    Run: func(cmd *cobra.Command, args []string) {
        if isRepositoriesList {
            repositories := artifactory.GetRepositories()
            fmt.Println(output.AnyToString(repositories))
        }
        if len(args) > 0 {
            if len(notUsedSince) > 0 {
                repositories := artifactory.GetArtifactNotUsedSinceForRepository(args[0], notUsedSince)
                fmt.Println(output.AnyToString(repositories))
                return
            }
            if len(deleteArtifact) > 0 {
                result := artifactory.DeleteArtifact(args[0], deleteArtifact)
                if (result) { 
                    fmt.Println("INFO: Deleted") 
                } else {
                    fmt.Println("ERR: Unable to delete") 
                }
                return
            }
        } else {
            fmt.Println("WARN: At least one repository must be passed in input")
            return
        }
    },
}

func init() {
    artifactoryCmd.AddCommand(artifactoryRepositoryCmd)
    artifactoryRepositoryCmd.Flags().BoolVar(&isRepositoriesList, "list", false, `List of all repositories`)
    artifactoryRepositoryCmd.Flags().StringVar(&notUsedSince, "not-used-since", "", `List all the artifact of a repository not used by a specified date
Examples:
    - imp artifactory repository test-local --not-used-since 2023-03-14
    `)
    artifactoryRepositoryCmd.Flags().StringVar(&deleteArtifact, "delete-artifact", "", `Remove the artifact on a specified repository
Examples:
    - imp artifactory repository test-local --delete-artifact test/local/1.2.3/artifact.json
    `)
}
