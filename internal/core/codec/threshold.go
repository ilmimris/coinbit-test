package codec

import (
	"encoding/json"
	"fmt"

	"github.com/ilmimris/coinbit-test/internal/core/domain"
	"github.com/ilmimris/coinbit-test/internal/core/port/outbond/pubsub"
)

type thresholdCodec struct{}

func NewThresholdCodec() pubsub.GokaCodec {
	return &thresholdCodec{}
}

func (tc *thresholdCodec) Encode(value interface{}) ([]byte, error) {
	if _, isThreshold := value.(*domain.Threshold); !isThreshold {
		return nil, fmt.Errorf("codec requirees value *Threshold, got %T", value)
	}
	v := value.(*domain.Threshold)
	return json.Marshal(v)
}

func (tc *thresholdCodec) Decode(value []byte) (interface{}, error) {
	var (
		threshold domain.Threshold
		err       error
	)

	err = json.Unmarshal(value, &threshold)
	if err != nil {
		return nil, fmt.Errorf("error unmarshalling threshold: %v", err)
	}
	return &threshold, nil
}
