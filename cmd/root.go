package cmd

import (
	"os"
	"tinygit/cmd/porcelain"
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
	rootCmd.AddCommand(plumbing.CommitTreeCmd)
	rootCmd.AddCommand(plumbing.ReadTreeCmd)
	rootCmd.AddCommand(plumbing.UpdateIndexCmd)
	rootCmd.AddCommand(plumbing.WriteTreeCmd)
}

func addPorcelain() {
	rootCmd.AddCommand(porcelain.InitCmd)
	rootCmd.AddCommand(porcelain.DiffCmd)
	rootCmd.AddCommand(porcelain.LogCmd)
}

func init() {
	addPlumbing()
	addPorcelain()
}
