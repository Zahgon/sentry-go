package sentryslog

import (
	"context"
	"log/slog"
)

func source(sourceKey string, r *slog.Record) slog.Attr {
	_ = "STUB: not implemented"
	return *new(slog.Attr)
}

type replaceAttrFn = func(groups []string, a slog.Attr) slog.Attr

func replaceAttrs(fn replaceAttrFn, groups []string, attrs ...slog.Attr) []slog.Attr {
	_ = "STUB: not implemented"
	return nil
}

func attrsToMap(attrs ...slog.Attr) map[string]any { _ = "STUB: not implemented"; return nil }

func extractError(attrs []slog.Attr) ([]slog.Attr, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func mergeAttrValues(values ...slog.Value) slog.Value {
	_ = "STUB: not implemented"
	return *new(slog.Value)
}

func groupValuesByKey(attrs []slog.Attr) map[string][]slog.Value {
	_ = "STUB: not implemented"
	return nil
}

func attrsToString(attrs ...slog.Attr) map[string]string { _ = "STUB: not implemented"; return nil }

func valueToString(v slog.Value) string { _ = "STUB: not implemented"; return "" }

func anyValueToString(v slog.Value) string { _ = "STUB: not implemented"; return "" }

func appendRecordAttrsToAttrs(attrs []slog.Attr, groups []string, record *slog.Record) []slog.Attr {
	_ = "STUB: not implemented"
	return nil
}

func removeEmptyAttrs(attrs []slog.Attr) []slog.Attr { _ = "STUB: not implemented"; return nil }

func contextExtractor(ctx context.Context, fns []func(ctx context.Context) []slog.Attr) []slog.Attr {
	_ = "STUB: not implemented"
	return nil
}

func appendAttrsToGroup(groups []string, actualAttrs []slog.Attr, newAttrs ...slog.Attr) []slog.Attr {
	_ = "STUB: not implemented"
	return nil
}

func toAnySlice(collection []slog.Attr) []any { _ = "STUB: not implemented"; return nil }

func uniqAttrs(attrs []slog.Attr) []slog.Attr { _ = "STUB: not implemented"; return nil }

func uniqByLast[T any, U comparable](collection []T, iteratee func(item T) U) []T {
	_ = "STUB: not implemented"
	return nil
}
