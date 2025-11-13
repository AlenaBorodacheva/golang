package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"strings"
)

func ReadFromCsv(filename string) ([]*Person, error) {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal("Ошибка открытия файла:", err)
		return nil, err
	}
	defer file.Close()
	reader := csv.NewReader(file)
	// Читаем заголовок (пропускаем)
	_, err = reader.Read()
	if err != nil && err != io.EOF {
		log.Fatal("Ошибка чтения заголовка:", err)
	}
	// Читаем данные
	res := []*Person{}
	for {
		record, err := reader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal("Ошибка чтения строки:", err)
			return nil, err
		}

		person, err := PersonFromCSVRecord(record)
		if err != nil {
			log.Printf("Ошибка парсинга строки: %v, строка: %v", err, record)
			continue
		}
		res = append(res, person)
	}

	return res, nil
}

func PersonFromCSVRecord(record []string) (*Person, error) {
	if len(record) != 4 {
		return nil, fmt.Errorf("ожидается 4 поля, получено %d", len(record))
	}
	name := record[0]
	age, err := strconv.Atoi(record[1])
	if err != nil {
		return nil, fmt.Errorf("не удалось прочитать Age: %v", err)
	}
	status, err := strconv.ParseBool(record[2])
	if err != nil {
		return nil, fmt.Errorf("не удалось прочитать Status: %v", err)
	}
	// Разбираем строку обратно в срез []int
	valuesStr := record[3]
	valuesStrSlice := strings.Split(valuesStr, ",")
	var values []int
	if valuesStr != "" { // Обработка пустой строки
		for _, s := range valuesStrSlice {
			s = strings.TrimSpace(s) // Убираем пробелы, если были
			v, err := strconv.Atoi(s)
			if err != nil {
				return nil, fmt.Errorf("не удалось прочитать элемент Values: %v", err)
			}
			values = append(values, v)
		}
	}
	return &Person{
		Name:   name,
		Age:    age,
		Status: status,
		Values: values,
	}, nil
}
