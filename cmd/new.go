/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"os"

	"github.com/fehlhabers/ensemble/data"
	"github.com/fehlhabers/ensemble/git"
	"github.com/spf13/cobra"
)

// newCmd represents the new command
var newCmd = &cobra.Command{
	Use:   "new",
	Short: "Starts new ensemble session",
	Long: `New creates a new branch both local and remote with the
	name provided as well as a prefix to identify the branch
	as an active ensemble session.

	`,
	Run: func(cmd *cobra.Command, args []string) {
		var (
			branch = fmt.Sprintf("%s%s", data.BranchPrefix, args[0])
		)

		fmt.Printf("Starting new ensemble session on <%s>...\n", branch)

		if err := git.Pull(); err != nil {
			fmt.Printf("WARNING! Not able to pull latest. Make sure your state is clean")
			os.Exit(1)
		}

		if err := git.NewBranch(args[0], "Starting new ensemble session..."); err != nil {
			fmt.Printf("WARNING! Unable to start new session")
			os.Exit(1)
		}

		fmt.Printf("Started new ensemble session <%s> !!! Good luck, team!\n", branch)
	},
}

func init() {
	rootCmd.AddCommand(newCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// newCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// newCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
