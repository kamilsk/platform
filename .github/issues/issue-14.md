---
code: null
id: MDU6SXNzdWU0NjQxOTIxMzY=
databaseId: 464192136
number: 14
url: https://github.com/kamilsk/platform/issues/14
title: "add pagination.Apply"
issueType: null
labels: []
assignees:
  - kamilsk
milestone: null
parent: null
issueFields: {}
state: CLOSED
stateReason: COMPLETED
duplicateOf: null
createdAt: 2019-07-04T10:30:53Z
updatedAt: 2019-07-09T10:46:16Z
lastEditedAt: 2019-07-09T09:01:10Z
closedAt: 2019-07-09T10:46:16Z
---

# add pagination.Apply

instead of

```go
	cnf := rest.PaginationConfiguration{PageKey: PageKey, PerPageKey: PerPageKey, PerPage: 30, PerPageMax: 100}
	return func() (string, gohttp.Handler) {
		return path, gohttp.HandlerFunc(func(rw gohttp.ResponseWriter, req *gohttp.Request) {

			...

			tmp := cnf
			perPageRaw := req.Form.Get(tmp.PerPageKey)
			if perPageRaw != "" {
				tmp.PerPage, _ = strconv.Atoi(perPageRaw)
			}
			if tmp.PerPage == 0 || tmp.PerPage > cnf.PerPageMax {
				tmp.PerPage = cnf.PerPage
			}
			page := 1
			pageRaw := req.Form.Get(tmp.PageKey)
			if pageRaw != "" {
				page, _ = strconv.Atoi(pageRaw)
			}
			if page < 1 {
				page = 1
			}
```

I want to

```go
			cnf := cnf.Apply(req.Form)
```

<!-- 2019-07-09T10:46Z https://github.com/kamilsk/platform/issues/14#issuecomment-509592290
fixed by 56bc263
-->
