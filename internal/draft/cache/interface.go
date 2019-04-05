package cache

import "context"

type Cache interface {
	Get(ctx context.Context, key string, value interface{}, options ...GetOption) error
	Set(ctx context.Context, key string, value interface{}, options ...SetOption) error
}

type Serializable interface {
	Marshaler
	Unmarshaler
}

type Marshaler interface{ Marshal() ([]byte, error) }

type Unmarshaler interface{ Unmarshal([]byte) error }
