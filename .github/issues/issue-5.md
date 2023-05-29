---
code: null
id: MDU6SXNzdWU0MjE1MDM2MjQ=
databaseId: 421503624
number: 5
url: https://github.com/kamilsk/platform/issues/5
title: "sort abstraction"
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
createdAt: 2019-03-15T12:50:37Z
updatedAt: 2020-03-16T19:36:50Z
lastEditedAt: null
closedAt: 2020-03-16T19:36:50Z
---

# sort abstraction

I want to add a high level abstraction above the built-in sort package to provide some advantages:

- `genome.Sort(what, how)` -> do sort
- `genome.Order(definitions ...OrderDefinition)` -> create composite sorting rule

Where `what` is a target to be sorted, `how` is a rule (simple or composite) about how a target will be sorted, `OrderDefinition` contains a transform mechanism (byName, byCount) and the sort direction (ASC or DESC) to apply sort.Reverse.

<!-- 2020-03-16T19:36Z https://github.com/kamilsk/platform/issues/5#issuecomment-599723609
moved to https://github.com/kamilsk/genome/issues/2
-->
