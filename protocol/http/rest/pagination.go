package rest

import (
	"fmt"
	"net/http"
	"net/url"
	"strconv"
	"strings"
)

type Pagination struct {
	First, Prev, Next, Last string
}

type PaginationConfig struct {
	PageKey      string
	PerPageKey   string
	Max, PerPage int
}

func Paginate(cnf PaginationConfig, total int, url *url.URL) (Pagination, error) {
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
	if limit > cnf.Max {
		limit = cnf.Max
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

func PaginationLink(header http.Header, pagination Pagination) {
	order := []string{"first", "prev", "next", "last"}
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
			hateoas = append(hateoas, fmt.Sprintf(`<%s>; rel="%s"`, v, k))
		}
	}
	if len(hateoas) > 0 {
		header.Add("Link", strings.Join(hateoas, ", "))
	}
}
