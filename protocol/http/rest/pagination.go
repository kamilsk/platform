package rest

import (
	"fmt"
	"net/http"
	"net/url"
	"strconv"
	"strings"
)

// Pagination holds information about page navigation.
//
// Deprecated: use go.octolab.org/toolkit/protocol/http/router/rest instead.
type Pagination struct{ First, Prev, Next, Last string }

// PaginationConfiguration holds a pagination configuration.
//
// Deprecated: use go.octolab.org/toolkit/protocol/http/router/rest instead.
type PaginationConfiguration struct {
	PageKey, PerPageKey string
	PerPage, PerPageMax int

	page, perPage int
}

// Apply fetches data from the values and applies it to a copy of the current configuration.
func (cnf PaginationConfiguration) Apply(values url.Values) PaginationConfiguration {
	cnf.page, _ = strconv.Atoi(values.Get(cnf.PageKey))
	if cnf.page < 1 {
		cnf.page = 1
	}
	limit, _ := strconv.Atoi(values.Get(cnf.PerPageKey))
	if limit == 0 {
		limit = cnf.PerPage
	}
	if limit > cnf.PerPageMax {
		limit = cnf.PerPageMax
	}
	cnf.perPage = 1
	if limit > 0 {
		cnf.perPage = limit
	}
	return cnf
}

// Limit returns the current pagination limit.
func (cnf *PaginationConfiguration) Limit() int { return cnf.perPage }

// Offset returns the current pagination offset.
func (cnf *PaginationConfiguration) Offset() int { return (cnf.page - 1) * cnf.Limit() }

// Paginate tries to build a pagination based on the configuration, total available items and current url.
func Paginate(url *url.URL, cnf PaginationConfiguration, total int) Pagination {
	var pagination Pagination

	limit := cnf.Limit()
	last := 1 + total/limit
	if total%limit == 0 {
		last--
	}

	q := url.Query()
	if limit != cnf.PerPage {
		q.Set(cnf.PerPageKey, strconv.Itoa(limit))
	}

	if cnf.page != 1 {
		q.Set(cnf.PageKey, "1")
		pagination.First = fmt.Sprintf("%s?%s", url.Path, q.Encode())

		q.Set(cnf.PageKey, strconv.Itoa(cnf.page-1))
		pagination.Prev = fmt.Sprintf("%s?%s", url.Path, q.Encode())
	}

	if cnf.page != last {
		q.Set(cnf.PageKey, strconv.Itoa(cnf.page+1))
		pagination.Next = fmt.Sprintf("%s?%s", url.Path, q.Encode())

		q.Set(cnf.PageKey, strconv.Itoa(last))
		pagination.Last = fmt.Sprintf("%s?%s", url.Path, q.Encode())
	}

	return pagination
}

// AddPaginationLink tries to add pagination links into the header.
//
// Deprecated: use go.octolab.org/toolkit/protocol/http/router/rest instead.
func AddPaginationLink(header http.Header, pagination Pagination) {
	order := [...]string{"first", "prev", "next", "last"}
	rel := map[string]*string{
		order[0]: &pagination.First,
		order[1]: &pagination.Prev,
		order[2]: &pagination.Next,
		order[3]: &pagination.Last,
	}
	hateoas := make([]string, 0, 4)
	for _, k := range order {
		v := *rel[k]
		if v != "" {
			hateoas = append(hateoas, fmt.Sprintf("<%s>; rel=%q", v, k))
		}
	}
	if len(hateoas) > 0 {
		header.Add("Link", strings.Join(hateoas, ", "))
	}
}
