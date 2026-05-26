// SPDX-License-Identifier: Apache-2.0
// Part of this code is derived from [github.com/johnbellone/grpc-middleware-sentry], licensed under the Apache 2.0 License.

package sentrygrpc

import (
	"context"
	"sync"

	"github.com/getsentry/sentry-go"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

const defaultClientOperationName = "rpc.client"

func hubFromClientContext(ctx context.Context) context.Context {
	_ = "STUB: not implemented"
	return *new(context.Context)
}

func createOrUpdateMetadata(ctx context.Context, span *sentry.Span) context.Context {
	_ = "STUB: not implemented"
	return *new(context.Context)
}

func finishSpan(span *sentry.Span, err error) { _ = "STUB: not implemented"; return }

func startClientSpan(ctx context.Context, method string) (context.Context, *sentry.Span) {
	_ = "STUB: not implemented"
	return *new(context.Context), nil
}

func UnaryClientInterceptor() grpc.UnaryClientInterceptor {
	_ = "STUB: not implemented"
	return *new(grpc.UnaryClientInterceptor)
}

func StreamClientInterceptor() grpc.StreamClientInterceptor {
	_ = "STUB: not implemented"
	return *new(grpc.StreamClientInterceptor)
}

type sentryClientStream struct {
	grpc.ClientStream
	serverStreams bool
	span          *sentry.Span
	stopMonitor   func() bool
	finishOnce    sync.Once
}

func (s *sentryClientStream) Header() (metadata.MD, error) {
	_ = "STUB: not implemented"
	return *new(metadata.MD), nil
}

func (s *sentryClientStream) SendMsg(m any) error { _ = "STUB: not implemented"; return nil }

func (s *sentryClientStream) RecvMsg(m any) error { _ = "STUB: not implemented"; return nil }

func (s *sentryClientStream) finish(err error) { _ = "STUB: not implemented"; return }
