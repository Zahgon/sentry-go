// SPDX-License-Identifier: Apache-2.0
// Part of this code is derived from [github.com/johnbellone/grpc-middleware-sentry], licensed under the Apache 2.0 License.

package sentrygrpc

import (
	"context"
	"time"

	"github.com/getsentry/sentry-go"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
)

const (
	sdkIdentifier              = "sentry.go.grpc"
	defaultServerOperationName = "rpc.server"
	internalServerErrorMessage = "internal server error"
)

type ServerOptions struct {
	// Repanic determines whether the application should re-panic after recovery.
	Repanic bool

	// WaitForDelivery determines if the interceptor should block until events are sent to Sentry.
	WaitForDelivery bool

	// Timeout sets the maximum duration for Sentry event delivery.
	Timeout time.Duration
}

func (o *ServerOptions) setDefaults() { _ = "STUB: not implemented"; return }

func recoverWithSentry(ctx context.Context, hub *sentry.Hub, o ServerOptions, onRecover func()) {
	_ = "STUB: not implemented"
	return
}

func hubFromServerContext(ctx context.Context) *sentry.Hub { _ = "STUB: not implemented"; return nil }

func traceHeadersFromContext(ctx context.Context) (metadata.MD, string, string) {
	_ = "STUB: not implemented"
	return *new(metadata.MD), "", ""
}

func startServerTransaction(ctx context.Context, fullMethod string) (context.Context, *sentry.Hub, *sentry.Span) {
	_ = "STUB: not implemented"
	return *new(context.Context), nil, nil
}

func setRPCStatus(span *sentry.Span, err error) { _ = "STUB: not implemented"; return }

func grpcStatusCode(err error) codes.Code { _ = "STUB: not implemented"; return *new(codes.Code) }

func UnaryServerInterceptor(opts ServerOptions) grpc.UnaryServerInterceptor {
	_ = "STUB: not implemented"
	return *new(grpc.UnaryServerInterceptor)
}

// StreamServerInterceptor provides Sentry integration for streaming gRPC calls.
func StreamServerInterceptor(opts ServerOptions) grpc.StreamServerInterceptor {
	_ = "STUB: not implemented"
	return *new(grpc.StreamServerInterceptor)
}

func getFirstHeader(md metadata.MD, key string) string { _ = "STUB: not implemented"; return "" }

func setScopeMetadata(hub *sentry.Hub, method string, md metadata.MD) {
	_ = "STUB: not implemented"
	return
}

func metadataToContext(md metadata.MD) map[string]any { _ = "STUB: not implemented"; return nil }

// parseGRPCMethod parses a gRPC full method name and returns the span name, service, and method components.
//
// It expects the format "/service/method" and parsing is compatible with:
// https://github.com/grpc/grpc-go/blob/v1.79.3/internal/grpcutil/method.go#L28
//
// Returns the original string as name and empty service/method if the format is invalid.
func parseGRPCMethod(fullMethod string) (name, service, method string) {
	_ = "STUB: not implemented"
	return "", "", ""
}

// wrapServerStream wraps a grpc.ServerStream, allowing you to inject a custom context.
func wrapServerStream(ctx context.Context, ss grpc.ServerStream) grpc.ServerStream {
	_ = "STUB: not implemented"
	return *new(grpc.ServerStream)
}

// wrappedServerStream is a wrapper around grpc.ServerStream that overrides the Context method.
type wrappedServerStream struct {
	grpc.ServerStream
	ctx context.Context
}

// Context returns the custom context for the stream.
func (w *wrappedServerStream) Context() context.Context {
	_ = "STUB: not implemented"
	return *new(context.Context)
}

var codeToSpanStatus = map[codes.Code]sentry.SpanStatus{
	codes.OK:                 sentry.SpanStatusOK,
	codes.Canceled:           sentry.SpanStatusCanceled,
	codes.Unknown:            sentry.SpanStatusUnknown,
	codes.InvalidArgument:    sentry.SpanStatusInvalidArgument,
	codes.DeadlineExceeded:   sentry.SpanStatusDeadlineExceeded,
	codes.NotFound:           sentry.SpanStatusNotFound,
	codes.AlreadyExists:      sentry.SpanStatusAlreadyExists,
	codes.PermissionDenied:   sentry.SpanStatusPermissionDenied,
	codes.ResourceExhausted:  sentry.SpanStatusResourceExhausted,
	codes.FailedPrecondition: sentry.SpanStatusFailedPrecondition,
	codes.Aborted:            sentry.SpanStatusAborted,
	codes.OutOfRange:         sentry.SpanStatusOutOfRange,
	codes.Unimplemented:      sentry.SpanStatusUnimplemented,
	codes.Internal:           sentry.SpanStatusInternalError,
	codes.Unavailable:        sentry.SpanStatusUnavailable,
	codes.DataLoss:           sentry.SpanStatusDataLoss,
	codes.Unauthenticated:    sentry.SpanStatusUnauthenticated,
}

func toSpanStatus(code codes.Code) sentry.SpanStatus {
	_ = "STUB: not implemented"
	return *new(sentry.SpanStatus)
}
