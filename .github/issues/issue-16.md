---
code: null
id: MDU6SXNzdWU0NjU4MjE1MzA=
databaseId: 465821530
number: 16
url: https://github.com/kamilsk/platform/issues/16
title: "sla middleware"
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
createdAt: 2019-07-09T14:41:35Z
updatedAt: 2019-07-10T07:18:44Z
lastEditedAt: 2019-07-09T14:43:13Z
closedAt: 2019-07-10T07:18:37Z
---

# sla middleware

```go
func SLA(defaults time.Duration, correction float64, handler gohttp.Handler) gohttp.Handler {
	return http.HandlerFunc(func(resp gohttp.ResponseWriter, req *gohttp.Request) {
		ctx, cancel := context.WithTimeout(req.Context(),
			http.Header(req.Header).Timeout(defaults, correction))
		defer cancel()

		handler.ServeHTTP(resp, req.WithContext(ctx))
	})
}
```
