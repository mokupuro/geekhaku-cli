/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"

	"github.com/manifoldco/promptui"
	"github.com/mokupuro/geekhaku-cli/utils"

	"github.com/spf13/cobra"
)

// geekhakuCmd represents the geekhaku command
var geekhakuCmd = &cobra.Command{
	Use:   "geekhaku",
	Short: "Print the ascii art of geekhaku",
	Run: func(cmd *cobra.Command, args []string) {
		selctType()
	},
}

func init() {
	rootCmd.AddCommand(geekhakuCmd)
}

func selctType() {

	prompt := promptui.Select{
		Label:     "What format do you want to display?",
		Items:     []string{"Default", "Kanji", "Shadow", "Moon", "3D"},
		CursorPos: 0,
	}

	idx, result, err := prompt.Run() //入力を受け取る

	if err != nil {
		fmt.Printf("Prompt failed %v\n", err)
		return
	}

	fmt.Printf("You choose no.%d %q\n", idx, result)

	switch result {

	case "Default":
		utils.PrintAAFromTxt("aa/geekhaku.txt")
	case "Kanji":
		utils.PrintAAFromTxt("aa/geekhaku_kanji.txt")
	case "Shadow":
		utils.PrintAAFromTxt("aa/geekhaku_shadow.txt")
	case "Moon":
		utils.PrintAAFromTxt("aa/geekhaku_moon.txt")
	case "3D":
		utils.PrintAAFromTxt("aa/geekhaku_3d.txt")
	}

}
