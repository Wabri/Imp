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

// hooksCmd represents the hook command
var hooksCmd = &cobra.Command{
	Use:   "hooks",
	Short: "Get list of gitlab project hooks",
	Long: `Get list of gitlab project hooks`,
    Run: func(cmd *cobra.Command, args []string) {
        id, err := strconv.Atoi(args[0])
        if err != nil {
            fmt.Printf("Something is wrong with the input: %v\n", args[0])
            os.Exit(1)
        }
        fmt.Println(gitlab.GetProjectHooksById(id)) 
    },
}

func init() {
	projectCmd.AddCommand(hooksCmd)
}
