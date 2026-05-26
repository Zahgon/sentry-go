package protocol

import (
	"errors"

	"github.com/getsentry/sentry-go/internal/ratelimit"
)

var errNoSerializableItems = errors.New("item container contains no serializable items")

type ItemContainer struct {
	items    []TelemetryItem
	category ratelimit.Category
}

// NewItemContainer constructs a batched envelope producer from buffered telemetry items.
func NewItemContainer(category ratelimit.Category, items []TelemetryItem) ItemContainer {
	_ = "STUB: not implemented"
	return *new(ItemContainer)
}

func (b ItemContainer) marshalPayload() ([]byte, int, error) {
	_ = "STUB: not implemented"
	return nil, 0, nil
}

func (b ItemContainer) newEnvelopeItem(itemCount int, payload []byte) (*EnvelopeItem, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (b ItemContainer) ToEnvelopeItem() (*EnvelopeItem, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (b ItemContainer) ToEnvelope(header *EnvelopeHeader) (*Envelope, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (b ItemContainer) GetCategory() ratelimit.Category {
	_ = "STUB: not implemented"
	return *new(ratelimit.Category)
}
func (ItemContainer) GetEventID() string   { _ = "STUB: not implemented"; return "" }
func (ItemContainer) GetSdkInfo() *SdkInfo { _ = "STUB: not implemented"; return nil }
func (ItemContainer) GetDynamicSamplingContext() map[string]string {
	_ = "STUB: not implemented"
	return nil
}
