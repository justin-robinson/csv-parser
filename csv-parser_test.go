package csv_parser

import (
	"encoding/csv"
	"log"
	"os"
	"strings"
	"testing"
)

var rawText = `"name","height","position","team"
"Bill Murray","3.7 gophers","grounds keeper","Bushwood Country Club"
"Jimmy Johnson","6","jsadklfadklvx","Cahbois"`

func TestParseString(t *testing.T) {

	cases := []struct {
		in   string
		want [][]string
	}{
		{
			rawText,
			[][]string{
				[]string{"\"name\"", "\"height\"", "\"position\"", "\"team\""},
				[]string{"\"Bill Murray\"", "\"3.7 gophers\"", "\"grounds keeper\"", "\"Bushwood Country Club\""},
				[]string{"\"Jimmy Johnson\"", "\"6\"", "\"jsadklfadklvx\"", "\"Cahbois\""},
			},
		},
	}

	for _, c := range cases {
		parser := NewCsvParser(strings.NewReader(c.in))
		got := parser.Parse()
		for i, row := range got {
			for j, value := range row {
				if value != c.want[i][j] {
					t.Errorf("ParseString == %q, want %q", value, c.want[i][j])
				}
			}
		}
	}
}

func TestParseFile(t *testing.T) {

	// open file
	file, err := os.Open("sample.csv")

	// return err if we have one
	if err != nil {
		log.Fatal(err)
	}

	parser := NewCsvParser(file)

	got := parser.Parse()

	if err != nil {
		t.Errorf("ParseFile %q", err)
	}

	want := [][]string{
		[]string{"policyID", "statecode", "county", "eq_site_limit", "hu_site_limit", "fl_site_limit", "fr_site_limit", "tiv_2011", "tiv_2012", "eq_site_deductible", "hu_site_deductible", "fl_site_deductible", "fr_site_deductible", "point_latitude", "point_longitude", "line", "construction", "point_granularity"},
		[]string{"119736", "FL", "CLAY COUNTY", "498960", "498960", "498960", "498960", "498960", "792148.9", "0", "9979.2", "0", "0", "30.102261", "-81.711777", "Residential", "Masonry", "1"},
		[]string{"448094", "FL", "CLAY COUNTY", "1322376.3", "1322376.3", "1322376.3", "1322376.3", "1322376.3", "1438163.57", "0", "0", "0", "0", "30.063936", "-81.707664", "Residential", "Masonry", "3"},
		[]string{"206893", "FL", "CLAY COUNTY", "190724.4", "190724.4", "190724.4", "190724.4", "190724.4", "192476.78", "0", "0", "0", "0", "30.089579", "-81.700455", "Residential", "Wood", "1"},
		[]string{"333743", "FL", "CLAY COUNTY", "0", "79520.76", "0", "0", "79520.76", "86854.48", "0", "0", "0", "0", "30.063236", "-81.707703", "Residential", "Wood", "3"},
		[]string{"172534", "FL", "CLAY COUNTY", "0", "254281.5", "0", "254281.5", "254281.5", "246144.49", "0", "0", "0", "0", "30.060614", "-81.702675", "Residential", "Wood", "1"},
		[]string{"785275", "FL", "CLAY COUNTY", "0", "515035.62", "0", "0", "515035.62", "884419.17", "0", "0", "0", "0", "30.063236", "-81.707703", "Residential", "Masonry", "3"},
		[]string{"995932", "FL", "CLAY COUNTY", "0", "19260000", "0", "0", "19260000", "20610000", "0", "0", "0", "0", "30.102226", "-81.713882", "Commercial", "Reinforced Concrete", "1"},
		[]string{"223488", "FL", "CLAY COUNTY", "328500", "328500", "328500", "328500", "328500", "348374.25", "0", "16425", "0", "0", "30.102217", "-81.707146", "Residential", "Wood", "1"},
		[]string{"433512", "FL", "CLAY COUNTY", "315000", "315000", "315000", "315000", "315000", "265821.57", "0", "15750", "0", "0", "30.118774", "-81.704613", "Residential", "Wood", "1"},
		[]string{"142071", "FL", "CLAY COUNTY", "705600", "705600", "705600", "705600", "705600", "1010842.56", "14112", "35280", "0", "0", "30.100628", "-81.703751", "Residential", "Masonry", "1"},
		[]string{"253816", "FL", "CLAY COUNTY", "831498.3", "831498.3", "831498.3", "831498.3", "831498.3", "1117791.48", "0", "0", "0", "0", "30.10216", "-81.719444", "Residential", "Masonry", "1"},
		[]string{"894922", "FL", "CLAY COUNTY", "0", "24059.09", "0", "0", "24059.09", "33952.19", "0", "0", "0", "0", "30.095957", "-81.695099", "Residential", "Wood", "1"},
		[]string{"422834", "FL", "CLAY COUNTY", "0", "48115.94", "0", "0", "48115.94", "66755.39", "0", "0", "0", "0", "30.100073", "-81.739822", "Residential", "Wood", "1"},
		[]string{"582721", "FL", "CLAY COUNTY", "0", "28869.12", "0", "0", "28869.12", "42826.99", "0", "0", "0", "0", "30.09248", "-81.725167", "Residential", "Wood", "1"},
		[]string{"842700", "FL", "CLAY COUNTY", "0", "56135.64", "0", "0", "56135.64", "50656.8", "0", "0", "0", "0", "30.101356", "-81.726248", "Residential", "Wood", "1"},
		[]string{"874333", "FL", "CLAY COUNTY", "0", "48115.94", "0", "0", "48115.94", "67905.07", "0", "0", "0", "0", "30.113743", "-81.727463", "Residential", "Wood", "1"},
		[]string{"580146", "FL", "CLAY COUNTY", "0", "48115.94", "0", "0", "48115.94", "66938.9", "0", "0", "0", "0", "30.121655", "-81.732391", "Residential", "Wood", "3"},
		[]string{"456149", "FL", "CLAY COUNTY", "0", "80192.49", "0", "0", "80192.49", "86421.04", "0", "0", "0", "0", "30.109537", "-81.741661", "Residential", "Wood", "1"},
		[]string{"767862", "FL", "CLAY COUNTY", "0", "48115.94", "0", "0", "48115.94", "73798.5", "0", "0", "0", "0", "30.11824", "-81.745335", "Residential", "Wood", "3"},
		[]string{"353022", "FL", "CLAY COUNTY", "0", "60946.79", "0", "0", "60946.79", "62467.29", "0", "0", "0", "0", "30.065799", "-81.717416", "Residential", "Wood", "1"},
		[]string{"367814", "FL", "CLAY COUNTY", "0", "28869.12", "0", "0", "28869.12", "42727.74", "0", "0", "0", "0", "30.082993", "-81.710581", "Residential", "Wood", "1"},
		[]string{"671392", "FL", "CLAY COUNTY", "0", "13410000", "0", "0", "13410000", "11700000", "0", "0", "0", "0", "30.091921", "-81.711929", "Commercial", "Reinforced Concrete", "3"},
		[]string{"772887", "FL", "CLAY COUNTY", "0", "1669113.93", "0", "0", "1669113.93", "2099127.76", "0", "0", "0", "0", "30.117352", "-81.711884", "Residential", "Masonry", "1"},
		[]string{"983122", "FL", "CLAY COUNTY", "0", "179562.23", "0", "0", "179562.23", "211372.57", "0", "0", "0", "0", "30.095783", "-81.713181", "Residential", "Wood", "3"},
		[]string{"934215", "FL", "CLAY COUNTY", "0", "177744.16", "0", "0", "177744.16", "157171.16", "0", "0", "0", "0", "30.110518", "-81.727478", "Residential", "Wood", "1"},
		[]string{"385951", "FL", "CLAY COUNTY", "0", "17757.58", "0", "0", "17757.58", "16948.72", "0", "0", "0", "0", "30.10288", "-81.705719", "Residential", "Wood", "1"},
		[]string{"716332", "FL", "CLAY COUNTY", "0", "130129.87", "0", "0", "130129.87", "101758.43", "0", "0", "0", "0", "30.068468", "-81.71624", "Residential", "Wood", "1"},
		[]string{"751262", "FL", "CLAY COUNTY", "0", "42854.77", "0", "0", "42854.77", "63592.88", "0", "0", "0", "0", "30.068468", "-81.71624", "Residential", "Wood", "1"},
		[]string{"633663", "FL", "CLAY COUNTY", "0", "785.58", "0", "0", "785.58", "662.18", "0", "0", "0", "0", "30.068468", "-81.71624", "Residential", "Wood", "1"},
		[]string{"105851", "FL", "CLAY COUNTY", "0", "170361.91", "0", "0", "170361.91", "177176.38", "0", "0", "0", "0", "30.068468", "-81.71624", "Residential", "Wood", "1"},
		[]string{"710400", "FL", "CLAY COUNTY", "0", "1430.89", "0", "0", "1430.89", "1861.41", "0", "0", "0", "0", "30.068468", "-81.71624", "Residential", "Wood", "1"},
		[]string{"703001", "FL", "CLAY COUNTY", "0", "129913.27", "0", "0", "129913.27", "101692.86", "0", "0", "0", "0", "30.079785", "-81.706865", "Residential", "Wood", "4"},
		[]string{"352792", "FL", "CLAY COUNTY", "0", "366285.62", "0", "0", "366285.62", "507164.19", "0", "0", "0", "0", "30.08012", "-81.718452", "Residential", "Masonry", "1"},
		[]string{"717603", "FL", "CLAY COUNTY", "0", "22512.61", "0", "0", "22512.61", "28637.17", "0", "0", "0", "0", "30.08012", "-81.718452", "Residential", "Wood", "1"},
		[]string{"937659", "FL", "SUWANNEE COUNTY", "0", "9246.6", "0", "9246.6", "9246.6", "10880.22", "0", "0", "0", "0", "29.959805", "-82.926659", "Residential", "Wood", "3"},
		[]string{"294022", "FL", "SUWANNEE COUNTY", "0", "96164.64", "0", "0", "96164.64", "69357.78", "0", "0", "0", "0", "29.959805", "-82.926659", "Residential", "Wood", "3"},
		[]string{"410500", "FL", "SUWANNEE COUNTY", "0", "11095.92", "0", "0", "11095.92", "12737.89", "0", "0", "0", "0", "29.959805", "-82.926659", "Residential", "Wood", "3"},
		[]string{"524433", "FL", "SUWANNEE COUNTY", "218475", "218475", "218475", "218475", "218475", "199030.29", "0", "4369.5", "0", "0", "29.962601", "-82.926155", "Residential", "Wood", "3"},
		[]string{"779298", "FL", "SUWANNEE COUNTY", "1400904", "1400904", "1400904", "1400904", "1400904", "1772984.1", "0", "0", "0", "0", "29.962601", "-82.926155", "Residential", "Masonry", "3"},
		[]string{"491831", "FL", "SUWANNEE COUNTY", "4365", "4365", "4365", "4365", "4365", "4438.05", "0", "87.3", "0", "0", "29.962601", "-82.926155", "Residential", "Wood", "3"},
		[]string{"814637", "FL", "SUWANNEE COUNTY", "4365", "4365", "4365", "4365", "4365", "6095.72", "0", "87.3", "0", "0", "29.962601", "-82.926155", "Residential", "Wood", "3"},
		[]string{"737515", "FL", "SUWANNEE COUNTY", "39789", "39789", "39789", "39789", "39789", "58106.58", "0", "0", "0", "0", "29.962601", "-82.926155", "Residential", "Wood", "3"},
		[]string{"222653", "FL", "SUWANNEE COUNTY", "24867", "24867", "24867", "24867", "24867", "18969.79", "0", "0", "0", "0", "29.962601", "-82.926155", "Residential", "Wood", "3"},
		[]string{"788543", "FL", "SUWANNEE COUNTY", "213876", "213876", "213876", "213876", "213876", "261435.18", "0", "0", "0", "0", "29.962601", "-82.926155", "Residential", "Wood", "3"},
		[]string{"691681", "FL", "SUWANNEE COUNTY", "69435", "69435", "69435", "69435", "69435", "93674.34", "0", "1388.7", "0", "0", "29.960735", "-82.92542", "Residential", "Wood", "3"},
		[]string{"368807", "FL", "SUWANNEE COUNTY", "14922", "14922", "14922", "14922", "14922", "12333.03", "0", "0", "0", "0", "29.960735", "-82.92542", "Residential", "Wood", "3"},
		[]string{"174002", "FL", "SUWANNEE COUNTY", "165546", "165546", "165546", "165546", "165546", "239134.51", "0", "0", "0", "0", "29.963396", "-82.916763", "Residential", "Wood", "1"},
		[]string{"198760", "FL", "SUWANNEE COUNTY", "72837", "72837", "72837", "72837", "72837", "86637.86", "0", "0", "0", "0", "29.963396", "-82.916763", "Residential", "Wood", "1"},
		[]string{"831395", "FL", "SUWANNEE COUNTY", "72837", "72837", "72837", "72837", "72837", "98147.86", "0", "0", "0", "0", "29.963396", "-82.916763", "Residential", "Wood", "1"},
	}

	for i, row := range got {
		for j, value := range row {
			if value != want[i][j] {
				t.Errorf("ParseFile == %q, want %q", value, want[i][j])
			}
		}
	}
}

func BenchmarkParseString(b *testing.B) {

	parser := NewCsvParser(strings.NewReader(rawText))

	for i := 0; i < b.N; i++ {
		parser.Parse()
	}
}

func BenchmarkParseFile(b *testing.B) {

	// open file
	file, err := os.Open("sample.csv")

	// return err if we have one
	if err != nil {
		log.Fatal(err)
	}

	parser := NewCsvParser(file)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		parser.Parse()
	}
}

func BenchmarkGoParseFile(b *testing.B) {
	file, err := os.Open("sample_large.csv")

	if err != nil {
		log.Fatal(err)
	}

	r := csv.NewReader(file)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		r.ReadAll()
	}
}
