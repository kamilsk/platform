---
code: null
id: MDU6SXNzdWU0MTcyMjA0NDQ=
databaseId: 417220444
number: 4
url: https://github.com/kamilsk/platform/issues/4
title: "add sequence generator"
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
createdAt: 2019-03-05T10:17:18Z
updatedAt: 2019-03-06T19:25:58Z
lastEditedAt: null
closedAt: 2019-03-06T19:25:58Z
---

# add sequence generator

```go
type generator int64

func (generator *generator) Next() int64 {
	return atomic.AddInt64((*int64)(generator), 1)
}
```
