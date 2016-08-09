package main

import (
	"errors"
	"fmt"
	"image"
	"os"
	"path/filepath"
	"strconv"
	"strings"

	"github.com/disintegration/imaging"
)

func main() {

	args := os.Args[1:]

	if len(args) != 5 {
		err := errors.New("Invalid arguments, example: ./resize-images assets/images/ sourceName.jpg destinationName.jpg 245 137")
		if err != nil {
			panic(err)
		}
	}

	searchDir := args[0]
	fileSource := args[1]
	fileDestination := args[2]

	fileWidth, e := strconv.Atoi(args[3])
	if e != nil {
		panic(e)
	}

	fileHeight, e := strconv.Atoi(args[4])
	if e != nil {
		panic(e)
	}

	fmt.Println("Searching directory:", searchDir)
	fmt.Println("File source:", fileSource)
	fmt.Println("File destination:", fileSource)
	fmt.Println("File size conversion to:", strconv.Itoa(fileWidth)+"x"+strconv.Itoa(fileHeight))

	fileList := []string{}
	err := filepath.Walk(searchDir, func(path string, f os.FileInfo, err error) error {
		if err != nil {
			panic(err)
		}
		index := strings.Index(path, fileSource)
		if index != -1 {
			fileList = append(fileList, path)
		}
		return nil
	})
	if err != nil {
		panic(err)
	}

	if len(fileList) == 0 {
		fmt.Println("No files found.")
	}

	for _, file := range fileList {
		img, err := imaging.Open(file)
		if err != nil {
			panic(err)
		}
		dstImage := imaging.Resize(img, fileWidth, fileHeight, imaging.NearestNeighbor)
		newPath := strings.Replace(file, fileSource, fileDestination, -1)
		saveFile(dstImage, newPath)
	}
}

func saveFile(file *image.NRGBA, path string) {
	fmt.Println("Saving file:", path)
	err := imaging.Save(file, path)
	if err != nil {
		panic(err)
	}
}
