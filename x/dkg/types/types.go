package types

import (
	"encoding/json"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

type Counter struct {
	Count uint64 `json:"count"`
}

func (c Counter) MustMarshalBinaryBare() []byte {
	bz, err := json.Marshal(c)
	if err != nil {
		panic(err) // handle the error according to your use case
	}
	return sdk.MustSortJSON(bz)
}

func (c *Counter) MustUnmarshalBinaryBare(bz []byte) {
	if err := json.Unmarshal(bz, c); err != nil {
		panic(err) // handle the error according to your use case
	}
}
