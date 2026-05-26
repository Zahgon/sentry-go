package sentry

import (
	"context"
	"sync"
	"time"
)

const (
	batchSize           = 100
	defaultBatchTimeout = 5 * time.Second
)

type batchProcessor[T any] struct {
	sendBatch    func([]T)
	itemCh       chan T
	flushCh      chan chan struct{}
	cancel       context.CancelFunc
	wg           sync.WaitGroup
	startOnce    sync.Once
	shutdownOnce sync.Once
	batchTimeout time.Duration
}

func newBatchProcessor[T any](sendBatch func([]T)) *batchProcessor[T] {
	_ = "STUB: not implemented"
	return nil
}

// WithBatchTimeout sets a custom batch timeout for the processor.
// This is useful for testing or when different timing behavior is needed.
func (p *batchProcessor[T]) WithBatchTimeout(timeout time.Duration) *batchProcessor[T] {
	_ = "STUB: not implemented"
	return nil
}

func (p *batchProcessor[T]) Send(item T) bool { _ = "STUB: not implemented"; return false }

func (p *batchProcessor[T]) Start() { _ = "STUB: not implemented"; return }

//nolint:gosec // G118: cancel is stored in p.cancel and called in Shutdown()

func (p *batchProcessor[T]) Flush(timeout <-chan struct{}) { _ = "STUB: not implemented"; return }

func (p *batchProcessor[T]) Shutdown() { _ = "STUB: not implemented"; return }

func (p *batchProcessor[T]) run(ctx context.Context) { _ = "STUB: not implemented"; return }
