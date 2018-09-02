package main

import (
	"fmt"
	"golang.org/x/text/collate"
	"golang.org/x/text/language"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

func usage() {
	fmt.Print(`Sort vietnamese words
Usage: vnsort <input file> [<output file>]
If output file is not set, the sorted result will be overwrite to input file.
`)
	os.Exit(1)
}

func main() {
	if len(os.Args) < 2 {
		usage()
	}
	inputFile := os.Args[1]
	outputFile := inputFile
	if len(os.Args) >= 3 {
		outputFile = os.Args[2]
	}
	b, err := ioutil.ReadFile(inputFile)
	if err != nil {
		log.Fatalln(err)
	}
	var words []string
	for _, w := range strings.Split(string(b), "\n") {
		if len(w) > 0 {
			words = append(words, w)
		}
	}

	vnm := collate.New(language.Vietnamese)
	vnm.SortStrings(words)

	err = ioutil.WriteFile(outputFile, []byte(strings.Join(words, "\n")), 0777)
	if err != nil {
		log.Fatalln(err)
	}
	log.Printf("Sorted %d lines", len(words))
}
