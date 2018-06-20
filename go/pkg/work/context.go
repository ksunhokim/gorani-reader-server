package work

import (
	"context"
	"errors"
)

type contextKey struct {
	name string
}

func (k *contextKey) String() string {
	return "sunho/gorani-reader/work context value " + k.name
}

var (
	ErrNoSuchValue = errors.New("work: no such value in context")
)

var (
	jobKey = contextKey{"job"}
	resKey = contextKey{"result"}
)

func GetJob(ctx context.Context) (Job, error) {
	if job, ok := ctx.Value(jobKey).(Job); ok {
		return job, nil
	}
	return Job{}, ErrNoSuchValue
}

func SetJob(ctx context.Context, job Job) context.Context {
	return context.WithValue(ctx, jobKey, job)
}

func GetResult(ctx context.Context) (Result, error) {
	if res, ok := ctx.Value(resKey).(Result); ok {
		return res, nil
	}
	return Result{}, ErrNoSuchValue
}

func SetResult(ctx context.Context, res Result) context.Context {
	return context.WithValue(ctx, resKey, res)
}
