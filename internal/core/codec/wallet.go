package codec

import (
	"encoding/json"
	"fmt"

	"github.com/ilmimris/coinbit-test/internal/core/domain"
	"github.com/ilmimris/coinbit-test/internal/core/port/outbond/pubsub"
)

type walletCodec struct{}

func NewWalletCodec() pubsub.GokaCodec {
	return &walletCodec{}
}

func (wc *walletCodec) Encode(value interface{}) ([]byte, error) {
	if _, isWallet := value.(*domain.Wallet); !isWallet {
		return nil, fmt.Errorf("codec requirees value *domain.Wallet, got %T", value)
	}
	v := value.(*domain.Wallet)
	return json.Marshal(v)
}

func (wc *walletCodec) Decode(value []byte) (interface{}, error) {
	var (
		wallet domain.Wallet
		err    error
	)

	err = json.Unmarshal(value, &wallet)
	if err != nil {
		return nil, fmt.Errorf("error unmarshalling wallet: %v", err)
	}

	return &wallet, nil
}
