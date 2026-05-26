package sentry

// logBatchProcessor batches logs and sends them to Sentry.
type logBatchProcessor struct {
	*batchProcessor[Log]
}

func newLogBatchProcessor(client *Client) *logBatchProcessor { _ = "STUB: not implemented"; return nil }

func (p *logBatchProcessor) Send(log *Log) bool { _ = "STUB: not implemented"; return false }
