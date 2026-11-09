package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/vinay03/chalk"
)

func checkerr(err error) {
	if err != nil {
		panic(err)
	}
}

const COL_WIDTH = 20

// FIX NAME OVERFLOW ISSUE
func serveDir(dir string) {
	f, err := os.OpenFile(dir, os.O_RDONLY, 0666)
	checkerr(err)
	finfo, err := f.Readdir(0)
	checkerr(err)

	var builder strings.Builder

	for _, entry := range finfo {
		var Name string
		if entry.IsDir() {
			Name = chalk.Green().Sprintf(entry.Name())
			fmt.Fprintf(&builder, "%-*s", COL_WIDTH+9, Name)
		} else {
			Name = entry.Name()
			fmt.Fprintf(&builder, "%-*s", COL_WIDTH, Name)
		}
		builder.WriteString(strconv.FormatInt(entry.Size(), 10))
		builder.WriteString("\n")

	}
	fmt.Println(builder.String())
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
