//go:generate echo $PWD - $GOPACKAGE - $GOFILE
//go:generate mockgen -package $GOPACKAGE -destination mock_server_test.go github.com/kamilsk/platform/protocol Server
package protocol_test

import (
	"context"
	"errors"
	"testing"
	"time"

	"github.com/golang/mock/gomock"
	_ "github.com/golang/mock/mockgen/model"
	"github.com/stretchr/testify/assert"

	. "github.com/kamilsk/platform/protocol"
)

func TestRun(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	tests := []struct {
		name     string
		server   func() Server
		shutdown func() Shutdown
		expected func(assert.TestingT, error, ...interface{}) bool
	}{
		{
			"listen and serve error",
			func() Server {
				mock := NewMockServer(ctrl)
				mock.EXPECT().ListenAndServe().Return(errors.New("listen and serve"))
				return mock
			},
			func() Shutdown { return nil },
			assert.Error,
		},
		{
			"listen and serve panic",
			func() Server {
				mock := NewMockServer(ctrl)
				mock.EXPECT().ListenAndServe().Do(func() { panic("bad server provider") })
				return mock
			},
			func() Shutdown { return nil },
			assert.Error,
		},
		{
			"shutdown error",
			func() Server {
				mock := NewMockServer(ctrl)
				mock.EXPECT().ListenAndServe().Do(func() { time.Sleep(100 * time.Millisecond) }).Return(nil).AnyTimes()
				mock.EXPECT().Shutdown(context.Background()).Return(errors.New("shutdown"))
				return mock
			},
			func() Shutdown {
				ch := make(Shutdown, 1)
				ch <- Callback{Context: context.Background(), Result: make(chan error, 1)}
				return ch
			},
			assert.Error,
		},
		{
			"graceful shutdown",
			func() Server {
				mock := NewMockServer(ctrl)
				mock.EXPECT().
					ListenAndServe().
					Do(func() { time.Sleep(100 * time.Millisecond) }).
					Return(errors.New("unexpected")).
					AnyTimes()
				mock.EXPECT().Shutdown(context.Background()).Return(nil)
				return mock
			},
			func() Shutdown {
				ch := make(Shutdown, 1)
				ch <- Callback{Context: context.Background(), Result: make(chan error, 1)}
				return ch
			},
			assert.NoError,
		},
	}
	for _, test := range tests {
		tc := test
		t.Run(test.name, func(t *testing.T) {
			tc.expected(t, Run(tc.server(), tc.shutdown()))
		})
	}
}
