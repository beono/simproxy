package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"testing"
)

func Test_Example(t *testing.T) {
	tests := []struct {
		name        string
		url         string
		expected    string
		contentType string
	}{
		{
			name:        "Must proxy",
			url:         "/oauth2/v3/certs",
			expected:    "true",
			contentType: "application/json; charset=UTF-8",
		},
		{
			name:        "Must not proxy",
			url:         "/hello",
			expected:    "false",
			contentType: "text/plain; charset=utf-8",
		},
	}

	go main()

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req, err := http.NewRequest(http.MethodGet, "http://localhost:8080"+tt.url, nil)
			if err != nil {
				t.Fatalf("unexpected error: %s", err)
			}
			resp, err := http.DefaultClient.Do(req)
			if err != nil {
				t.Fatalf("unexpected error: %s", err)
			}

			if resp.Header.Get("x-proxy") != tt.expected {
				t.Errorf(`unexpected value for "x-proxy", got: %q`, resp.Header.Get("content-type"))
			}

			if resp.Header.Get("content-type") != tt.contentType {
				t.Errorf(`unexpected value for "content-type", got: %q`, resp.Header.Get("content-type"))
			}

			if t.Failed() {
				defer resp.Body.Close()
				body, _ := ioutil.ReadAll(resp.Body)
				fmt.Printf("Reponse body: %s", body)
			}
		})
	}
}
