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
