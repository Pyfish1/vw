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

const COL_WIDTH = 30
const BUFFER = 1

func shorten(entry string) string {
	const MAX_LENGTH = COL_WIDTH - BUFFER
	if len(entry) > MAX_LENGTH {
		return entry[:MAX_LENGTH-2] + ".."
	} else {
		return entry
	}
}

func serveDir(dir string) {
	f, err := os.OpenFile(dir, os.O_RDONLY, 0666)
	checkerr(err)
	finfo, err := f.Readdir(0)
	checkerr(err)

	var builder strings.Builder

	fmt.Fprintf(&builder, "%-*s", COL_WIDTH, "Name")
	builder.WriteString("Size")
	builder.WriteString("\n")
	fmt.Fprintf(&builder, strings.Repeat("-", COL_WIDTH+6)+"\n")

	for _, entry := range finfo {
		var Name string
		if entry.IsDir() {
			Name = shorten(entry.Name())
			Name = chalk.Green().Sprintf(Name)
			fmt.Fprintf(&builder, "%-*s", COL_WIDTH+9, Name)
		} else {
			Name = shorten(entry.Name())
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
