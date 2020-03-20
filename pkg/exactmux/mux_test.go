package exactmux

import (
	"fmt"
	"net/http"
	"strings"
	"testing"
)

func Test_calculateWeight(t *testing.T) {
	tests := []struct {
		name    string
		pattern string
		want    int
	}{
		{"zero for root", "/", 0},
		{"one for root + subpath", "/catalog", 1},
		{"two for root + subpath (with trailing slash)", "/catalog/", 2},
		{"three for root + hierarchical subpath", "/catalog/1", 3},
		{"four for root + hierarchical subpath (with trailing slash)", "/catalog/1/", 4},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := calculateWeight(tt.pattern); got != tt.want {
				t.Errorf("calculateWeight() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_paramRoutesMatch(t *testing.T) {
	entry := paramsMuxEntry{pathParts: strings.Split("/hello/{id}/comm/{commId}/", "/"),
		handler: http.HandlerFunc(func(http.ResponseWriter, *http.Request) {
			fmt.Println("200 OK")
		}),
	}
	if ok := paramRoutesMatch(entry, "/hello/123/comm/456/"); ok != true {
		t.Error("fucck")
	}
}
