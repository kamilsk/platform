---
code: null
id: MDU6SXNzdWU0NTIzMzk5NDU=
databaseId: 452339945
number: 9
url: https://github.com/kamilsk/platform/issues/9
title: "pack method for chi"
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
createdAt: 2019-06-05T06:49:22Z
updatedAt: 2019-06-08T09:50:29Z
lastEditedAt: 2019-06-06T20:33:11Z
closedAt: 2019-06-08T09:50:29Z
---

# pack method for chi

```go
import (
	gohttp "net/http"

	"github.com/go-chi/chi"
	"go.avito.ru/swat/service-swaha/internal/protocol/http"
)

type Handler func() (string, string, gohttp.HandlerFunc)

func Pack(method string, handler http.Handler, placeholders ...string) Handler {
	if len(placeholders)%2 != 0 {
		panic("passed placeholders must contain from => to list")
	}
	return func() (string, string, gohttp.HandlerFunc) {
		p, h := handler()
		return method, p, func(rw gohttp.ResponseWriter, req *gohttp.Request) {
			if len(placeholders) > 0 {
				q := req.URL.Query()
				for i, steps := 0, len(placeholders)/2; i <= steps; i += 2 {
					from, to := placeholders[i], placeholders[i+1]
					q.Set(to, chi.URLParam(req, from))
				}
				req.URL.RawQuery = q.Encode()
			}
			h.ServeHTTP(rw, req)
		}
	}
}
```
