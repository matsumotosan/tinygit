package plumbing

import (
	"os"
	"crypto/sha1"
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
	if _, err := os.Stat(path); os.IsNotExist(err) {
		if err := os.MkdirAll(path, os.ModePerm); err != nil {
			log.Fatal(err)
		}
		fmt.Printf("Create %s", path)
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
