---
code: null
id: MDU6SXNzdWU0NTM4MTY4NTg=
databaseId: 453816858
number: 10
url: https://github.com/kamilsk/platform/issues/10
title: "hateoas pagination"
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
createdAt: 2019-06-08T17:31:08Z
updatedAt: 2019-06-08T17:36:42Z
lastEditedAt: null
closedAt: 2019-06-08T17:36:42Z
---

# hateoas pagination

```
type Pagination struct {
	First, Prev, Next, Last string
}

type PaginationConfig struct {
	PageKey      string
	PerPageKey   string
	Max, PerPage int
}
```
