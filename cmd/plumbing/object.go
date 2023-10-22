package plumbing

import (
	"os"
	"crypto/sha1"
	"path/filepath"
	"fmt"
	"log"
)

type Object struct {
	Type string
	Size int
	Content []byte
}

func NewObject(objectType string, content []byte) Object {
	return Object{
		Type: objectType,
		Size: len(content),
		Content: content,
	}
}

func GetObject(sha1 string) Object {
	content, err := os.ReadFile(Sha2path(sha1))
	if err != nil {
		log.Fatal(err)
	}

	return NewObject("blob", content)
}

func (o Object) PrettyPrint() {
	switch o.Type {
	case "commit":
		fmt.Println("not implemented")
	case "tree":
		fmt.Println("not implemented")
	case "blob":
		fmt.Println(string(o.Content))
	case "tag":
		fmt.Println("not implemented")
	}
}

func (o Object) Save(path string) {
	if _, err := os.Stat(filepath.Dir(path)); os.IsNotExist(err) {
		if err := os.MkdirAll(filepath.Dir(path), os.ModePerm); err != nil {
			log.Fatal(err)
		}
	}

	if err := os.WriteFile(path, []byte(o.Content), 0666); err != nil {
		log.Fatal(err)
	}
}

func (o Object) Sha1Hash() string {
	var data []byte
	header := []byte(fmt.Sprintf("%s %d", o.Type, o.Size))
	data = append(data, header...)
	data = append(data, '\x00')
	data = append(data, o.Content...)
	return fmt.Sprintf("%x", sha1.Sum(data))
}

func Sha2path(sha1 string) string {
	return filepath.Join(".tinygit/objects", sha1[:2], sha1[2:])
}
