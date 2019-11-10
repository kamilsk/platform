package protobuf

import (
	"time"

	"github.com/golang/protobuf/ptypes"
	"github.com/golang/protobuf/ptypes/timestamp"
	"github.com/pkg/errors"
)

// Time converts protobuf Timestamp into built-in Time.
// Deprecated: use go.octolab.org/toolkit/protocol/protobuf.Time.
func Time(ts *timestamp.Timestamp) *time.Time {
	if ts == nil {
		return nil
	}
	tp, err := ptypes.Timestamp(ts)
	if err != nil {
		panic(errors.Wrapf(err, "converting %#v into time.Time", *ts))
	}
	return &tp
}

// Time converts built-in Time into protobuf Timestamp.
// Deprecated: use go.octolab.org/toolkit/protocol/protobuf.Timestamp.
func Timestamp(tp *time.Time) *timestamp.Timestamp {
	if tp == nil {
		return nil
	}
	ts, err := ptypes.TimestampProto(*tp)
	if err != nil {
		panic(errors.Wrapf(err, "converting %#v into google.protobuf.Timestamp", *tp))
	}
	return ts
}
