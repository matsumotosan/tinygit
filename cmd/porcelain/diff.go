package porcelain

import (
	"fmt"

	"github.com/spf13/cobra"
)

var DiffCmd = &cobra.Command{
	Use:   "diff",
	Short: "Compare file with index.",
	Long:  `Compare file with index.`,
	Run:   diffRun,
}

func diffRun(cmd *cobra.Command, args []string) {
	fmt.Println("tinygit diff called")
}

func init() {
}
