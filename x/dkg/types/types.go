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

type TimeoutData struct {
	Round uint64 `json:"round"`
	Id string `json:"id"`
	Timeout uint64 `json:"timeout"`
	Start uint64 `json:"start"`
}

func (t TimeoutData) MustMarshalBinaryBare() []byte {
	bz, err := json.Marshal(t)
	if err != nil {
		panic(err) // handle the error according to your use case
	}
	return sdk.MustSortJSON(bz)
}

func (t *TimeoutData) MustUnmarshalBinaryBare(bz []byte) {
	if err := json.Unmarshal(bz, t); err != nil {
		panic(err) // handle the error according to your use case
	}
}