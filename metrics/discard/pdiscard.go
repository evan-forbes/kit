// Package discard provides a no-op metrics backend.
package discard

import (
	"github.com/go-kit/kit/metrics"
	"github.com/prometheus/client_golang/prometheus"
)

type pcounter struct{}

// NewCounter returns a new no-op counter.
func NewPCounter() metrics.PCounter { return pcounter{} }

// With implements Counter.
func (c pcounter) With(labelValues ...string) metrics.Counter { return c }

func (c pcounter) WithP(labelValues ...string) metrics.PCounter { return c }

func (c pcounter) Collector() prometheus.Collector { return nil }

// Add implements Counter.
func (c pcounter) Add(delta float64) {}

type pgauge struct{}

// pNewGauge returns a new no-op pgauge.
func NewPGauge() metrics.PGauge { return pgauge{} }

// With implements pGauge.
func (g pgauge) With(labelValues ...string) metrics.Gauge { return g }

// Set implements pGauge.
func (g pgauge) Set(value float64) {}

// Add implements metrics.pGauge.
func (g pgauge) Add(delta float64) {}

func (g pgauge) WithP(labelValues ...string) metrics.PGauge { return g }

func (g pgauge) Collector() prometheus.Collector { return nil }

type phistogram struct{}

// NewHistogram returns a new no-op histogram.
func NewPHistogram() metrics.PHistogram { return phistogram{} }

// With implements Histogram.
func (h phistogram) With(labelValues ...string) metrics.Histogram { return h }

func (h phistogram) WithP(labelValues ...string) metrics.PHistogram { return h }

func (h phistogram) Collector() prometheus.Collector { return nil }

// Observe implements phistogram.
func (h phistogram) Observe(value float64) {}
