package csv

import (
	"bufio"
	"encoding/csv"
	"os"
	"strings"
)

type CsvReader struct {
	reader *csv.Reader
}

func NewCsvReader(fileContent *string) *CsvReader {
	if fileContent == nil {
		return nil
	}

	reader := csv.NewReader(strings.NewReader(*fileContent))

	return &CsvReader{reader: reader}
}

func (r *CsvReader) ReadAll() ([][]string, error) {
	return r.reader.ReadAll()
}

func NewCsvBufioReader(csvPath string) (*CsvReader, error) {

	csvFile, err := os.Open(csvPath)
	if err != nil {
		return nil, err
	}
	reader := csv.NewReader(bufio.NewReader(csvFile))

	return &CsvReader{reader: reader}, nil
}

func NewCsvReaderAll(csvFile string) ([][]string, error) {

	reader, err := NewCsvBufioReader(csvFile)
	if err != nil {
		return nil, err
	}
	return reader.ReadAll()

}
