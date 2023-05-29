---
code: null
id: MDU6SXNzdWU0NTYxMDU5Nzk=
databaseId: 456105979
number: 11
url: https://github.com/kamilsk/platform/issues/11
title: "middleware for rest"
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
createdAt: 2019-06-14T07:41:52Z
updatedAt: 2019-06-20T14:40:48Z
lastEditedAt: null
closedAt: 2019-06-20T14:40:48Z
---

# middleware for rest

```go
type Middleware func(Handler) Handler

type Configuration struct { Handlers, Middlewares }

type Option func(RouterConfiguration)

func WithMiddlewares(...Middleware) Option {}

func WithHandlers(...Handler) Option {}
```
