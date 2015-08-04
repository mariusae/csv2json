// Csv2json reads CSV records from standard in, and outputs a JSON
// array of map entries to standard out. Csv2json assumes the first
// CSV record is the header.
package main // import "marius.ae/csv2json"

import (
	"bufio"
	"encoding/csv"
	"encoding/json"
	"io"
	"log"
	"os"
)

func main() {
	log.SetFlags(0)
	log.SetPrefix("csv2json: ")

	in := bufio.NewReader(os.Stdin)
	r := csv.NewReader(in)
	r.FieldsPerRecord = -1

	header, err := r.Read()
	if err != nil {
		log.Fatal(err)
	}

	var records []map[string]string

	for {

		fields, err := r.Read()
		if err == io.EOF {
			break
		} else if err != nil {
			log.Fatal(err)
		}

		record := make(map[string]string)
		for i := range fields {
			if i == len(header) {
				break
			}

			record[header[i]] = fields[i]
		}

		records = append(records, record)
	}

	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	e := json.NewEncoder(out)

	if err := e.Encode(records); err != nil {
		log.Fatal(err)
	}
}