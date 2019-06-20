package rest

import (
	"fmt"
	"net/http"
	"net/url"
	"strconv"
	"strings"
)

// Pagination holds information about page navigation.
type Pagination struct {
	First, Prev, Next, Last string
}

// PaginationConfiguration holds a pagination configuration.
type PaginationConfiguration struct {
	PageKey, PerPageKey string
	PerPage, PerPageMax int
}

// Paginate tries to build a pagination based on the configuration, total available items and current url.
func Paginate(cnf PaginationConfiguration, url *url.URL, total int) (Pagination, error) {
	var pagination Pagination

	q := url.Query()

	current := 1
	page, found := q[cnf.PageKey]
	if found && len(page) > 0 {
		var err error
		current, err = strconv.Atoi(page[0])
		if err != nil {
			return pagination, err
		}
	}

	limit := cnf.PerPage
	perPage, found := q[cnf.PerPageKey]
	if found && len(perPage) > 0 {
		var err error
		limit, err = strconv.Atoi(perPage[0])
		if err != nil {
			return pagination, err
		}
	}
	if cnf.PerPageMax > 0 && limit > cnf.PerPageMax {
		limit = cnf.PerPageMax
		q.Set(cnf.PerPageKey, strconv.Itoa(limit))
	}

	last := 1 + total/limit
	if total%limit == 0 {
		last--
	}

	if current != 1 {
		q.Set(cnf.PageKey, "1")
		pagination.First = fmt.Sprintf("%s?%s", url.Path, q.Encode())

		q.Set(cnf.PageKey, strconv.Itoa(current-1))
		pagination.Prev = fmt.Sprintf("%s?%s", url.Path, q.Encode())
	}

	if current != last {
		q.Set(cnf.PageKey, strconv.Itoa(current+1))
		pagination.Next = fmt.Sprintf("%s?%s", url.Path, q.Encode())

		q.Set(cnf.PageKey, strconv.Itoa(last))
		pagination.Last = fmt.Sprintf("%s?%s", url.Path, q.Encode())
	}

	return pagination, nil
}

// AddPaginationLink tries to add pagination links into the header.
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
