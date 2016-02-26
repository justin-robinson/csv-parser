package csv_parser

import (
	"bufio"
	"strings"
)

type CsvParser struct {
	scanner *bufio.Scanner
}


func (c *CsvParser) Parse() ([][]string) {

	// where are parsed lines will be stored
	parsed := [][]string{}

	// add each line to our parsed array
	for c.scanner.Scan() {
		parsed = append(parsed, strings.Split(c.scanner.Text(), ","))
	}

	return parsed
}
