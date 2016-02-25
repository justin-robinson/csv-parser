package csv_parser

import (
	"bufio"
	"os"
	"strings"
)

type CsvParser struct {}


func parse(scanner *bufio.Scanner) ([][]string) {

	// where are parsed lines will be stored
	parsed := [][]string{}

	// add each line to our parsed array
	for scanner.Scan() {
		parsed = append(parsed, strings.Split(scanner.Text(), ","))
	}

	return parsed
}

func (c *CsvParser) ParseFile(fileName string) ([][]string, error) {

	// open file
	file, err:= os.Open(fileName)

	// return err if we have one
	if err != nil {
		return [][]string{}, err
	}

	return parse(bufio.NewScanner(file)), nil
}


func (c *CsvParser) ParseString(input string) ([][]string) {

	return parse(bufio.NewScanner(strings.NewReader(input)))
}
