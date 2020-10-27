// Package scanner_stub реализует интерфейс Scanner и выдает закодированные данные.
package scanner_stub

type ScannerStub struct {
}

func (s *ScannerStub) Scan(url string, depth int) (map[string]string, error) {
	data := make(map[string]string)
	data["google.com"] = "google"
	data["ya.ru"] = "ya"
	return data, nil
}
