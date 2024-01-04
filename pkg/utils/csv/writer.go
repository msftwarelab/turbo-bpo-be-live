package csv

import (
	"bufio"
	"bytes"
	"encoding/csv"
)

type CsvWriter struct {
	buff   *bytes.Buffer
	writer *csv.Writer
}

func NewCsvWriter() *CsvWriter {
	var buff bytes.Buffer
	writer := csv.NewWriter(bufio.NewWriter(&buff))

	return &CsvWriter{buff: &buff, writer: writer}
}

func (w *CsvWriter) WriteAll(lines [][]string) ([]byte, error) {
	err := w.writer.WriteAll(lines)
	if err != nil {
		return nil, err
	}

	w.writer.Flush()

	return w.buff.Bytes(), nil
}
