/*
Copyright © 2023 Gabriele Puliti <gabriele.puliti+github@proton.me>
*/
package cmd

import (
	"fmt"
	"imp/addons/gitlab"
	"os"
	"strconv"

	"github.com/spf13/cobra"
)

// projectCmd represents the project command
var projectCmd = &cobra.Command{
	Use:   "project",
	Short: "Get github project by id",
	Long: `Get github project by id`,
	Run: func(cmd *cobra.Command, args []string) {
        id, err := strconv.Atoi(args[0])
        if err != nil {
            fmt.Printf("Something is wrong with the input: %v\n", args[0])
            os.Exit(1)
        }
        fmt.Println(gitlab.GetProjectById(id)) 
	},
}

func init() {
	gitlabCmd.AddCommand(projectCmd)
}
