package cmd

import (
	"os"
	tinyinit "tinygit/cmd/init"
	"tinygit/cmd/diff"
	"tinygit/cmd/plumbing"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "tinygit",
	Short: "A tiny reimplementation of git in Go.",
	Long: `A tiny reimplementation of git in Go.`,
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func addPlumbing() {
	rootCmd.AddCommand(plumbing.HashObjectCmd)
	rootCmd.AddCommand(plumbing.CatFileCmd)
}

func addPorcelain() {
	rootCmd.AddCommand(tinyinit.InitCmd)
	rootCmd.AddCommand(diff.DiffCmd)
}

func init() {
	addPlumbing()
	addPorcelain()
}
