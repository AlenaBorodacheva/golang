package writer

import (
	"bytes"
	"encoding/csv"
	"encoding/json"
	"log"
	"os"

	"github.com/yukithm/json2csv"
)

func WriteToCsv(filename string, data []byte) error {
	// --- Запись в CSV ---
	file, err := os.OpenFile(filename, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644)
	if err != nil {
		log.Fatal("Ошибка создания файла:", err)
		return err
	}
	defer file.Close()

	var x []map[string]interface{}

	// unMarshall json
	err1 := json.Unmarshal(data, &x)
	if err1 != nil {
		log.Fatal(err)
	}
	csv1, err := json2csv.JSON2CSV(x)
	if err != nil {
		log.Fatal(err)
	}

	// CSV bytes convert & writing...
	b := &bytes.Buffer{}
	wr := json2csv.NewCSVWriter(b)
	err = wr.WriteCSV(csv1)
	if err != nil {
		log.Fatal(err)
	}
	wr.Flush()
	got := b.String()

	//Following line prints CSV
	println(got)

	writer := csv.NewWriter(file)
	defer writer.Flush() // Убедиться, что буфер записан

	if _, err = file.WriteString(got); err != nil {
		panic(err)
	}

	return nil
}
