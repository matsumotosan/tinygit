package plumbing

import (
	"io"
	"os"
	"fmt"
	"path/filepath"

	"github.com/spf13/cobra"
)

var (
	content []byte
	objectType string
	write bool
	stdin bool
	stdinPaths bool
	path string
)

const (
	commit = iota
	tree
	blob
	tag
)

var HashObjectCmd = &cobra.Command{
	Use:   "hash-object",
	Short: "Compute object ID value for an object",
	Long:  `Compute object ID value for an object`,
	// Args:  cobra.MinimumNArgs(1),
	Run:   hashObjectRun,
}

func hashObjectRun(cmd *cobra.Command, args []string) {
	if stdin {
		content, _ = io.ReadAll(os.Stdin)
	} else if stdinPaths {
		fmt.Println("not implemented")
	} else {
		fmt.Println("not implemented")
	}

	object := NewObject(objectType, content)
	sha1 := object.Sha1Hash()

	if write {
		path = sha2path(sha1)
		object.Save(path)
	}

	fmt.Println(sha1)
}

func sha2path(sha1 string) string {
	return filepath.Join(".tinygit/objects", sha1[:2], sha1[2:])
}

func init() {
	HashObjectCmd.Flags().StringVarP(&objectType, "type", "t", "blob", "type of object to be created")
	HashObjectCmd.Flags().BoolVarP(&write, "write", "w", true, "write object to object database")
	HashObjectCmd.Flags().BoolVar(&stdin, "stdin", false, "read object from standard input instead of from a file")
	HashObjectCmd.Flags().BoolVar(&stdin, "stdin-paths", false, "read file names from standard input")
	HashObjectCmd.Flags().StringVar(&path, "path", "", "hash object as if it were located at the given path")
}
