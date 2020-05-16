package main

import (
	"flag"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"path/filepath"
	"strings"
)

var err error
var path = flag.String("p", "", "provide abs file path")
var filename = flag.String("f", "", "provide file name")
var truncate = flag.String("t", "", "provide ending value to search and remove")

func main() {
	flag.Parse()

	fileName := *filename
	filePath := *path

	if filePath == "" {
		filePath, err = filepath.Abs(fileName + "/" + fileName)
		if err != nil {
			fmt.Println("provide file")
			return
		}
	}

	cleanFile := filePath + fileName
	fmt.Println("clean file: " + cleanFile)

	input, err := ioutil.ReadFile(cleanFile)
	if err != io.EOF && err != nil {
		fmt.Println("Failed " + err.Error())
		return
	}

	data := string(input)

	if len(data) == 0 {
		fmt.Println("empty")
		return
	}
	index := strings.LastIndex(data, *truncate)
	fmt.Println(index)
	if index > -1 {
		//fmt.Println(data[index:])
		data = data[:index]
	}
	err = ioutil.WriteFile(cleanFile, []byte(data), 0644)
	if err != nil {
		log.Fatalln(err)
	}
}
