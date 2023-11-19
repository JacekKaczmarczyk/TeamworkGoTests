package main

import (
	"customerdomains/customerimporter"
	"fmt"
	"io"
	"log"
	"os"
)

type FileReader struct{}

func (*FileReader) Open(fileName string) (file io.ReadCloser, err error) {
	return os.Open(fileName)
}

func main() {
	fileReader := &FileReader{}
	domians, err := customerimporter.CountAndSortEmailDomains(fileReader)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Print(domians)
}
