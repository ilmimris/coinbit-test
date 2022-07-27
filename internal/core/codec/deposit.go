package codec

import (
	"fmt"

	domain "github.com/ilmimris/coinbit-test/internal/core/domain/proto"
	"github.com/ilmimris/coinbit-test/internal/core/port/outbond/pubsub"
	"google.golang.org/protobuf/proto"
)

type depositCodec struct {
}

func NewDepositCodec() pubsub.GokaCodec {
	return &depositCodec{}
}

func (jc *depositCodec) Encode(value interface{}) ([]byte, error) {
	if _, isDeposit := value.(*domain.Wallet); !isDeposit {
		return nil, fmt.Errorf("Codec requires value *domain.DepositWallet, got %T", value)
	}
	v := value.(*domain.Wallet)
	return proto.Marshal(v)
}

// Decodes a deposit from []byte to it's go representation.
func (jc *depositCodec) Decode(data []byte) (interface{}, error) {
	var (
		deposit domain.Wallet
		err     error
	)
	err = proto.Unmarshal(data, &deposit)
	if err != nil {
		return nil, fmt.Errorf("Error unmarshaling deposit: %v", err)
	}
	return &deposit, nil
}
