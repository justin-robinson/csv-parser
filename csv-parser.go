package csv_parser

import (
	"bufio"
	"os"
	"strings"
)

type CsvParser struct {
}

func (c *CsvParser) Parse(fileName string) ([][]string, error) {

	// where are parsed lines will be stored
	parsed := [][]string{}

	// open file
	file, err:= os.Open(fileName)

	// return err if we have one
	if err != nil {
		return [][]string{}, err
	}

	// create a scanner to scan the csv
	scanner := bufio.NewScanner(file)

	// add each line to our parsed array
	for i:=0; scanner.Scan(); i++ {
		parsed = insert(parsed, i, strings.Split(scanner.Text(), ","))
	}

	return parsed, nil
}

func insert(original [][]string, position int, value []string) [][]string {
	l := len(original)
	target := original

	// does the array have any capacity left?
	if cap(original) == l {

		// create a new array with more room to grow
		target = make([][]string, l+1, l+10)

		// copy values after the insertion point from original to target
		copy(target, original[:position])
	} else {
		// append an empty entry
		target = append(target, []string{})
	}

	// copy original values before the insertion point to the target,
	// this leaves the position open
	copy(target[position+1:], original[position:])

	// set the position to the new value
	target[position] = value

	return target
}
