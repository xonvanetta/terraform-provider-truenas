package http

import (
	"context"
	"errors"
)

//TODO: better package name with interface

type Client interface {
	Post(ctx context.Context, path string, body interface{}, v interface{}) error
	Get(ctx context.Context, path string, v interface{}) error
	Put(ctx context.Context, path string, body interface{}, v interface{}) error
	Delete(ctx context.Context, path string) error
}

var ErrNotImplemented = errors.New("not implemented")
