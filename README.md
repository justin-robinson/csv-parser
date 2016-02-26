# csv-parser

Don't use this. Use go's built in csv parser https://golang.org/pkg/encoding/csv/
HOWEVER this one is about 50x faster if you don't care as much about edge cases

```shell
go test -bench File
PASS
BenchmarkParseFile-4  	50000000	      22.1 ns/op
BenchmarkGoParseFile-4	 1000000	      1021 ns/op
ok  	github.com/justin-robinson/csv-parser	8.982s
```

handles comma delimited csv files and does not checked for escaped characters.

```go
package main

import (
	"fmt"
	"github.com/justin-robinson/csv-parser"
)

func main () {

    // open file
    file, err:= os.Open("sample.csv")

    // return err if we have one
    if err != nil {
        log.Fatal(err)
    }

    parser := CsvParser{
        bufio.NewScanner(file),
    }

    parsed := parser.Parse()
    
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
