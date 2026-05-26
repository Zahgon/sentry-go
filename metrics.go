package sentry

import (
	"context"
	"sync"

	"github.com/getsentry/sentry-go/attribute"
)

// Duration Units.
const (
	UnitNanosecond  = "nanosecond"
	UnitMicrosecond = "microsecond"
	UnitMillisecond = "millisecond"
	UnitSecond      = "second"
	UnitMinute      = "minute"
	UnitHour        = "hour"
	UnitDay         = "day"
	UnitWeek        = "week"
)

// Information Units.
const (
	UnitBit      = "bit"
	UnitByte     = "byte"
	UnitKilobyte = "kilobyte"
	UnitKibibyte = "kibibyte"
	UnitMegabyte = "megabyte"
	UnitMebibyte = "mebibyte"
	UnitGigabyte = "gigabyte"
	UnitGibibyte = "gibibyte"
	UnitTerabyte = "terabyte"
	UnitTebibyte = "tebibyte"
	UnitPetabyte = "petabyte"
	UnitPebibyte = "pebibyte"
	UnitExabyte  = "exabyte"
	UnitExbibyte = "exbibyte"
)

// Fraction Units.
const (
	UnitRatio   = "ratio"
	UnitPercent = "percent"
)

// NewMeter returns a new Meter. If there is no Client bound to the current hub, or if metrics are disabled,
// it returns a no-op Meter that discards all metrics.
func NewMeter(ctx context.Context) Meter { _ = "STUB: not implemented"; return *new(Meter) }

// build default attrs

type sentryMeter struct {
	ctx               context.Context
	hub               *Hub
	attributes        map[string]attribute.Value
	defaultAttributes map[string]attribute.Value
	mu                sync.RWMutex
}

func (m *sentryMeter) emit(ctx context.Context, metricType MetricType, name string, value MetricValue, unit string, attributes map[string]attribute.Value, customScope *Scope) {
	_ = "STUB: not implemented"
	return
}

// Pre-allocate with capacity hint to avoid map growth reallocations
// scope ~3 + call-specific ~5

// attribute precedence: default -> scope -> instance (from SetAttrs) -> entry-specific

// WithCtx returns a new Meter that uses the given context for trace/span association.
func (m *sentryMeter) WithCtx(ctx context.Context) Meter {
	_ = "STUB: not implemented"
	return *new(Meter)
}

func (m *sentryMeter) applyOptions(opts []MeterOption) *meterOptions {
	_ = "STUB: not implemented"
	return nil
}

// Count implements Meter.
func (m *sentryMeter) Count(name string, count int64, opts ...MeterOption) {
	_ = "STUB: not implemented"
	return
}

// Distribution implements Meter.
func (m *sentryMeter) Distribution(name string, sample float64, opts ...MeterOption) {
	_ = "STUB: not implemented"
	return
}

// Gauge implements Meter.
func (m *sentryMeter) Gauge(name string, value float64, opts ...MeterOption) {
	_ = "STUB: not implemented"
	return
}

// SetAttributes implements Meter.
func (m *sentryMeter) SetAttributes(attrs ...attribute.Builder) { _ = "STUB: not implemented"; return }

// noopMeter is a no-operation implementation of Meter.
// This is used when there is no client available in the context or when metrics are disabled.
type noopMeter struct{}

// WithCtx implements Meter.
func (n *noopMeter) WithCtx(_ context.Context) Meter {
	_ = "STUB: not implemented"

	// Count implements Meter.
	return *new(Meter)
}

func (n *noopMeter) Count(name string, _ int64, _ ...MeterOption) {
	_ = "STUB: not implemented"
	return
}

// Distribution implements Meter.
func (n *noopMeter) Distribution(name string, _ float64, _ ...MeterOption) {
	_ = "STUB: not implemented"
	return
}

// Gauge implements Meter.
func (n *noopMeter) Gauge(name string, _ float64, _ ...MeterOption) {
	_ = "STUB: not implemented"
	return
}

// SetAttributes implements Meter.
func (n *noopMeter) SetAttributes(_ ...attribute.Builder) { _ = "STUB: not implemented"; return }
