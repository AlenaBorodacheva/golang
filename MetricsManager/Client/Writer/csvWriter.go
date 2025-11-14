package writer

import (
	"bytes"
	"encoding/json"
	"log"
	"os"

	"github.com/yukithm/json2csv"
)

func WriteToCsv(filename string, data []byte) error {
	b := &bytes.Buffer{}
	wr := json2csv.NewCSVWriter(b)
	var x []map[string]interface{}

	// unMarshall json
	err := json.Unmarshal(data, &x)
	if err != nil {
		log.Fatal(err)
		return err
	}

	// convert json to CSV
	csv, err := json2csv.JSON2CSV(x)
	if err != nil {
		log.Fatal(err)
		return err
	}

	// CSV bytes convert & writing...
	err = wr.WriteCSV(csv)
	if err != nil {
		log.Fatal(err)
		return err
	}
	wr.Flush()
	got := b.String()

	//Following line prints CSV
	println(got)

	createFileAppendText(filename, got)

	return nil
}

func createFileAppendText(filename string, text string) {
	f, err := os.OpenFile(filename, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600)
	if err != nil {
		panic(err)
	}

	defer f.Close()

	if _, err = f.WriteString(text); err != nil {
		panic(err)
	}
}
