package rest_test

import (
	"bytes"
	"net/http"
	"net/url"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"

	. "github.com/kamilsk/platform/protocol/http/rest"
)

func TestPaginate(t *testing.T) {
	tests := []struct {
		name     string
		url      *url.URL
		cnf      PaginationConfig
		total    int
		expected Pagination
	}{
		{
			"first page",
			&url.URL{Path: "/test"},
			PaginationConfig{"page", "per_page", 10, 3}, 7,
			Pagination{Next: "/test?page=2", Last: "/test?page=3"},
		},
		{
			"middle page",
			&url.URL{Path: "/test", RawQuery: "page=2"},
			PaginationConfig{"page", "per_page", 10, 3}, 7,
			Pagination{First: "/test?page=1", Prev: "/test?page=1", Next: "/test?page=3", Last: "/test?page=3"},
		},
		{
			"last page",
			&url.URL{Path: "/test", RawQuery: "page=3"},
			PaginationConfig{"page", "per_page", 10, 3}, 7,
			Pagination{First: "/test?page=1", Prev: "/test?page=2"},
		},
		{
			"overflow",
			&url.URL{Path: "/test", RawQuery: "per_page=30"},
			PaginationConfig{"page", "per_page", 10, 3}, 7,
			Pagination{},
		},
	}
	for _, test := range tests {
		tc := test
		t.Run(test.name, func(t *testing.T) {
			pagination, err := Paginate(tc.cnf, tc.total, tc.url)
			assert.NoError(t, err)
			assert.Equal(t, tc.expected, pagination)
		})
	}
}

func TestPaginationLink(t *testing.T) {
	tests := []struct {
		name       string
		pagination Pagination
		expected   string
	}{
		{
			"first page",
			Pagination{Next: "page=2", Last: "page=5"},
			`Link: <page=2>; rel="next", <page=5>; rel="last"`,
		},
		{
			"last page",
			Pagination{First: "page=1", Prev: "page=4"},
			`Link: <page=1>; rel="first", <page=4>; rel="prev"`,
		},
		{
			"middle page",
			Pagination{First: "page=1", Prev: "page=2", Next: "page=4", Last: "page=5"},
			`Link: <page=1>; rel="first", <page=2>; rel="prev", <page=4>; rel="next", <page=5>; rel="last"`,
		},
		{
			"overflow",
			Pagination{},
			"",
		},
	}
	for _, test := range tests {
		tc := test
		t.Run(test.name, func(t *testing.T) {
			buf, header := bytes.NewBuffer(nil), http.Header{}
			PaginationLink(header, tc.pagination)
			assert.NoError(t, header.Write(buf))
			assert.Equal(t, tc.expected, strings.TrimRight(buf.String(), "\r\n"))
		})
	}
}
