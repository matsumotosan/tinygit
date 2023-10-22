package plumbing

import (
	"fmt"

	"github.com/spf13/cobra"
)

var (
	exists bool
	prettyPrint bool
	showType bool
	showSize bool
)

var CatFileCmd = &cobra.Command{
	Use:   "cat-file",
	Short: "A brief description of your command",
	Long:  `A brief description of your command`,
	Args:  cobra.MinimumNArgs(1),
	Run:   catFileRun,
}

func catFileRun(cmd *cobra.Command, args []string) {
	object := GetObject(args[0])
	if prettyPrint {
		object.PrettyPrint()
	}

	if showType {
		fmt.Println("type")
	}

	if showSize {
		fmt.Println("size")
	}
}

func init() {
	CatFileCmd.Flags().BoolVarP(&exists, "exists", "e", false, "check if <object> exists")
	CatFileCmd.Flags().BoolVarP(&prettyPrint, "pretty", "p", false, "pretty print <object> content")
	CatFileCmd.Flags().BoolVarP(&showType, "show-type", "t", false, "show object type")
	CatFileCmd.Flags().BoolVarP(&showSize, "show", "s", false, "show object size")
}
