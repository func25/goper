package development

import (
	"context"
	"sync"
	"time"
)

type Context struct{}

var (
	levelMtx sync.Mutex
	fast     = 500 * time.Millisecond
	medium   = 3 * time.Second
	slow     = 10 * time.Second
)

func GetValue[T any](ctx context.Context, key string) (T, bool) {
	res, ok := ctx.Value(key).(T)
	return res, ok
}

func (Context) SetFast(t time.Duration) {
	levelMtx.Lock()
	defer levelMtx.Unlock()

	fast = t
}

func (Context) SetMedium(t time.Duration) {
	levelMtx.Lock()
	defer levelMtx.Unlock()

	medium = t
}

func (Context) SetSlow(t time.Duration) {
	levelMtx.Lock()
	defer levelMtx.Unlock()

	slow = t
}

func (Context) Fast(ctx context.Context) (context.Context, context.CancelFunc) {
	levelMtx.Lock()
	defer levelMtx.Unlock()

	return context.WithTimeout(ctx, fast)
}

func (Context) Medium(ctx context.Context) (context.Context, context.CancelFunc) {
	levelMtx.Lock()
	defer levelMtx.Unlock()

	return context.WithTimeout(ctx, medium)
}

func (Context) Slow(ctx context.Context) (context.Context, context.CancelFunc) {
	levelMtx.Lock()
	defer levelMtx.Unlock()

	return context.WithTimeout(ctx, slow)

}
