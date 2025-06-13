package services

import (
	"bytes"
	"context"
	"encoding/json"
	"io"
	"mime/multipart"
	"net/http"
	"os"
	"time"
)

type Tradeline struct {
	Creditor string `json:"creditor"`
	Account  string `json:"account"`
	Status   string `json:"status"`
}

type ParsedReport struct {
	Tradelines []Tradeline `json:"tradelines"`
	Negatives  []Tradeline `json:"negatives"`
}

// ParseReport sends the PDF to a local parser microservice
func ParseReport(path string) (ParsedReport, error) {
	file, err := os.Open(path)
	if err != nil {
		return ParsedReport{}, err
	}
	defer file.Close()

	var body bytes.Buffer
	writer := multipart.NewWriter(&body)
	part, err := writer.CreateFormFile("file", path)
	if err != nil {
		return ParsedReport{}, err
	}
	if _, err := io.Copy(part, file); err != nil {
		return ParsedReport{}, err
	}
	writer.Close()

	url := os.Getenv("PARSER_URL")
	if url == "" {
		url = "http://localhost:5000/parse"
	}
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	req, err := http.NewRequestWithContext(ctx, "POST", url, &body)
	if err != nil {
		return ParsedReport{}, err
	}
	req.Header.Set("Content-Type", writer.FormDataContentType())
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return ParsedReport{}, err
	}
	defer resp.Body.Close()

	var parsed ParsedReport
	if err := json.NewDecoder(resp.Body).Decode(&parsed); err != nil {
		return ParsedReport{}, err
	}
	return parsed, nil
}
