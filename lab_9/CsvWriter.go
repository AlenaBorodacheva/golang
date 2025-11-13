package main

import (
	"encoding/csv"
	"errors"
	"log"
	"os"
	"strconv"
	"strings"
)

func WriteToCsv(filename string, person Person) error {
	// --- Запись в CSV ---
	emptyFile := false
	if _, err := os.Stat(filename); errors.Is(err, os.ErrNotExist) {
		emptyFile = true
	}

	file, err := os.OpenFile(filename, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644)
	if err != nil {
		log.Fatal("Ошибка создания файла:", err)
		return err
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush() // Убедиться, что буфер записан

	if emptyFile {
		// Записываем заголовки (опционально)
		header := []string{"Name", "Age", "Status", "Values"}
		if err := writer.Write(header); err != nil {
			log.Fatal("Ошибка записи заголовка:", err)
			return err
		}
	}

	// Записываем данные
	record := person.ToCSVRecord()
	if err := writer.Write(record); err != nil {
		log.Fatal("Ошибка записи данных:", err)
		return err
	}
	return nil
}

// строка — это удобное представление Person для CSV
func (p *Person) ToCSVRecord() []string {
	// Преобразуем срез []int в строку, разделённую запятыми
	valuesStr := make([]string, len(p.Values))
	for i, v := range p.Values {
		valuesStr[i] = strconv.Itoa(v)
	}
	return []string{
		p.Name,
		strconv.Itoa(p.Age),
		strconv.FormatBool(p.Status),
		strings.Join(valuesStr, ","), // Объединяем срез в строку
	}
}
