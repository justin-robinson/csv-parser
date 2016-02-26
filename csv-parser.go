package csv_parser

import (
	"bufio"
	"io"
	"strings"
)

type csvParser struct {
	scanner *bufio.Scanner
}


func (c *csvParser) Parse() ([][]string) {

	// where are parsed lines will be stored
	parsed := [][]string{}

	// add each line to our parsed array
	for c.scanner.Scan() {
		parsed = append(parsed, strings.Split(c.scanner.Text(), ","))
	}

	return parsed
}

func NewCsvParser(r io.Reader) *csvParser {
	return &csvParser{
		bufio.NewScanner(r),
	}
}
