package main

import (
	"go-hardcode/hw5/pkg/scanner_stub"
	"testing"
)

func Test_scanLinks(t *testing.T) {
	s := new(scanner_stub.ScannerStub)
	data, err := scanLinks(s, "https://go.dev/", 2)
	if err != nil {
		t.Fatal(err)
	}

	for k, v := range data {
		t.Logf("%s -> %s\n", k, v)
	}
}
