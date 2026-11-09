package main

import (
	"fmt"
	"os"

	"github.com/vinay03/chalk"
)

func checkerr(err error) {
	if err != nil {
		panic(err)
	}
}

func serveDir(dir string) {
	f, err := os.OpenFile(dir, os.O_RDONLY, 0666)
	checkerr(err)
	finfo, err := f.Readdir(0)
	checkerr(err)

	for _, entry := range finfo {
		if entry.IsDir() {
			fileName := chalk.Green().Sprintf(entry.Name())
			fmt.Println(fileName)
		} else {
			fileName := entry.Name()
			fmt.Println(fileName)
		}
	}
}

func main() {
	arguments := os.Args[1:]
	var dir string
	if len(arguments) > 0 {
		dir = arguments[0]
	} else {
		dir = "."
	}
	serveDir(dir)
}
