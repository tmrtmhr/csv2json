package main

// convert csv to json using header column as key
// USAGE: csv2json -delimiter="\t" < sample.tsv > sample.json

import (
	"encoding/csv"
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
)

var delimiter *string = flag.String("delimiter", ",", "specify separator (e.g. \"\\t\")")
var lazyQuote *bool = flag.Bool("lazyQuote", true, "allow lazyQuote")
var jsonSeq *bool = flag.Bool("jsonSeq", false, "output JSON text sequence format")

type JSON map[string]interface{}

func main() {
	flag.Parse()

	// Unquote Escaped Character (e.g. \t)
	rune, _, _, err := strconv.UnquoteChar(*delimiter, '"')
	if err != nil {
		log.Fatalf("Error: UnquoteChar fail: Input: '%s', Message: %v\n", *delimiter, err)
	}

	// create CSV reader from stdin
	r := csv.NewReader(os.Stdin)
	r.Comma = rune
	r.LazyQuotes = *lazyQuote

	results := []JSON{}

	// read header
	header, err := r.Read()
	if err == io.EOF {
		return
	}
	for {
		// read csv body
		rows, err := r.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf("Error: csv Read fail: %v\n", err)
		}

		jsonData := make(JSON)
		for i := range rows {
			jsonData[header[i]] = string(rows[i])
		}

		if *jsonSeq {
			// output immediately
			body, err := json.Marshal(jsonData)
			if err != nil {
				log.Fatalf("Error: json.Marshal fail: Input: %v, Message: %v", results, err)
			}
			fmt.Printf("%s\n", body)
		} else {
			results = append(results, jsonData)
		}
	}

	if !*jsonSeq {
		// Save the JSON file
		json, err := json.MarshalIndent(results, "", "  ")
		if err != nil {
			log.Fatalf("Error: json.Marshal fail: Input: %v, Message: %v", results, err)
		}

		fmt.Printf("%s\n", json)
	}
}
