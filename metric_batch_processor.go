package sentry

// metricBatchProcessor batches metrics and sends them to Sentry.
type metricBatchProcessor struct {
	*batchProcessor[Metric]
}

func newMetricBatchProcessor(client *Client) *metricBatchProcessor {
	_ = "STUB: not implemented"
	return nil
}

func (p *metricBatchProcessor) Send(metric *Metric) bool { _ = "STUB: not implemented"; return false }
