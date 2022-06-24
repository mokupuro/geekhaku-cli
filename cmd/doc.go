/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"github.com/mokupuro/geekhaku-cli/utils"

	"github.com/spf13/cobra"
)

// docCmd represents the doc command
var docCmd = &cobra.Command{
	Use:   "doc",
	Short: "Print the ascii art of doc",
	Run: func(cmd *cobra.Command, args []string) {
		utils.PrintAAFromTxt("./aa/doc.txt")
	},
}

func init() {
	rootCmd.AddCommand(docCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// docCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// docCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
