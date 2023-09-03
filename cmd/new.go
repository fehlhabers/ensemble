/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

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
		fmt.Printf("Got arguments: %+v\n", args)
		fmt.Println("new called")
		git.NewBranch(args[0], "Starting new ensemble session...")
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
