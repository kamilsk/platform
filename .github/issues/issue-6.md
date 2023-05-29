---
code: null
id: MDU6SXNzdWU0MjE1MDczMDI=
databaseId: 421507302
number: 6
url: https://github.com/kamilsk/platform/issues/6
title: "generator draft"
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
createdAt: 2019-03-15T13:00:08Z
updatedAt: 2020-03-16T19:33:40Z
lastEditedAt: 2019-03-15T13:00:48Z
closedAt: 2020-03-16T19:33:39Z
---

# generator draft

The concept:

```go
sequence := new(integer.generator).Since(10).Rhythm(func (int64) int64).Until(10^5)

for range sequence.Source() { ... } // math.MaxInt64 by default

sequence.Next() bool

sequence.Current() int64

sequence.FastForward(int64) bool
```

```go
uid := new(uuid.generator)

uid.Next() string ?
```

```go
random := new(rand.generator)

random.Next() ?
```

<!-- 2020-03-16T19:33Z https://github.com/kamilsk/platform/issues/6#issuecomment-599722243
moved to https://github.com/octolab/pkg/issues/7
-->
