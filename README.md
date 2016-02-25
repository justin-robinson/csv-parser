# csv-parser

Don't use this. Use go's built in csv parser https://golang.org/pkg/encoding/csv/

handles comma delimited csv files and does not checked for escaped characters

```golang
package main

import (
	"fmt"
	"github.com/justin-robinson/csv-parser"
)

func main () {

    // create a parser
	parser := csv_parser.CsvParser{}

    // err from os.Open
	parsed, err := parser.Parse("/some/file.csv")

    
	if err != nil {
		panic(err)
	}

    // print out header row
	for i:=0; i<len(parsed[0]); i++ {
		fmt.Println(parsed[0][i])
	}

    // print out all rows
	fmt.Println(parsed)
}
```

/some/file.csv
```csv
"name","height","position","team"
"Bill Murray","3.7 gophers","grounds keeper","Bushwood Country Club"
"Jimmy Johnson","6","jsadklfadklvx","Cahbois"
```
