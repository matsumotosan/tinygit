package init

import (
	"fmt"
	"log"
	"os"
	"path"
	"path/filepath"

	"github.com/spf13/cobra"
)

var (
	quiet bool
	bare bool
	initialBranch string
)

var InitCmd = &cobra.Command{
	Use:   "init",
	Short: "Initialize a Tinygit repository.",
	Long: `Initialize a Tinygit repository.`,
	Run: initRun,
}

func initRun(cmd *cobra.Command, args []string) {
	dir, err := filepath.Abs("./.tinygit/")
	if err != nil {
		log.Fatal(err)
	}

	_, err = os.Stat(dir)

	initHooks(dir)
	initInfo(dir)
	initObjects(dir)
	initRefs(dir)

	writeHEAD(dir)
	writeDefaultConfig(dir)
	writeDescription(dir)

	if !quiet {
		printMessage(filepath.Dir(dir), err)
	}
}

func printMessage(dir string, err error) {
	var msg string
	if os.IsNotExist(err) {
		msg = "Initialized a new tinygit repository at"
	} else {
		msg = "Reinitialized a tinygit repository at"
	}
	fmt.Println(msg, dir)

}

func initHooks(dir string) {
	if err := os.MkdirAll(path.Join(dir, "hooks"), 0755); err != nil {
		log.Fatal(err)
	}
}

func initInfo(dir string) {
	if err := os.MkdirAll(path.Join(dir, "info"), 0755); err != nil {
		log.Fatal(err)
	}
}

func initObjects(dir string) {
	if err := os.MkdirAll(path.Join(dir, "objects/info"), 0755); err != nil {
		log.Fatal(err)
	}

	if err := os.MkdirAll(path.Join(dir, "objects/pack"), 0755); err != nil {
		log.Fatal(err)
	}
}

func initRefs(dir string) {
	if err := os.MkdirAll(path.Join(dir, "refs", "heads"), 0755); err != nil {
		log.Fatal(err)
	}

	if err := os.MkdirAll(path.Join(dir, "refs", "tags"), 0755); err != nil {
		log.Fatal(err)
	}
}

func writeDescription(dir string) {
	if err := os.WriteFile(path.Join(dir, "description"), []byte("Unnamed repository; edit this file 'description' to name the repository."), 0666); err != nil {
		log.Fatal(err)
	}
}

func writeHEAD(dir string) {
	if err := os.WriteFile(path.Join(dir, "HEAD"), []byte("ref: refs/head/" + initialBranch), 0666); err != nil {
		log.Fatal(err)
	}
}

func init() {
	InitCmd.Flags().BoolVarP(&quiet, "quiet", "q", false, "shut up")
	InitCmd.Flags().BoolVar(&bare, "bare", false, "create bare repository")
	InitCmd.Flags().StringVarP(&initialBranch, "initial-branch", "b", "master", "override initial branch name")
}
