/*
PACKAGE DOCUMENTATION

package csv_parser
    import "."

    Package csv_parser provides an alternative to the built in csv library
    offering more speed at the cost of error checking and other
    flexibilities

FUNCTIONS

func NewCsvParser(r io.Reader) *csvParser
    NewCsvParser constructs a new csvParser struct from an io.Reader
*/

// Package csv_parser provides an alternative to the built in csv library offering
// more speed at the cost of error checking and other flexibilities
package csv_parser

import (
	"bufio"
	"io"
	"strings"
)

// Private csvParser struct
type csvParser struct {
	// scanner points to the csv data we are parsing
	scanner *bufio.Scanner
}

// Parse parses the scanner contents into a [][]string array containing
// the data.  Where the first index is the row and the second being the value
func (c *csvParser) Parse() ([][]string) {

	// where are parsed lines will be stored
	parsed := [][]string{}

	// add each line to our parsed array
	for c.scanner.Scan() {
		parsed = append(parsed, strings.Split(c.scanner.Text(), ","))
	}

	return parsed
}

// NewCsvParser constructs a new csvParser struct from an io.Reader
func NewCsvParser(r io.Reader) *csvParser {
	return &csvParser{
		bufio.NewScanner(r),
	}
}
