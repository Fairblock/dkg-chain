package types

import (
	"encoding/json"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/sirupsen/logrus"
	//"github.com/drand/kyber"
)

type Counter struct {
	Count uint64 `json:"count"`
}

type AddressList struct {
	Addresses []string `json:"addresses"`
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
	Round   uint64 `json:"round"`
	Id      string `json:"id"`
	Timeout uint64 `json:"timeout"`
	Start   uint64 `json:"start"`
}

type Faulters struct {
	FaultyList []uint64        `json:"faultyList"`
	Lookup     map[uint64]bool `json:"lookup"`
}

// Initialize the map

type MPKData struct {
	Pks map[uint64][]byte `json:"pks"`
	Id  string            `json:"id"`
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

func (m MPKData) MustMarshalBinaryBare() []byte {
	bz, err := json.Marshal(m)
	if err != nil {
		logrus.Info(err) // handle the error according to your use case
	}
	return sdk.MustSortJSON(bz)
}

func (m *MPKData) MustUnmarshalBinaryBare(bz []byte) {
	if err := json.Unmarshal(bz, m); err != nil {
		panic(err) // handle the error according to your use case
	}
}

func (m Faulters) MustMarshalBinaryBare() []byte {
	bz, err := json.Marshal(m)
	if err != nil {
		logrus.Info(err) // handle the error according to your use case
	}
	return sdk.MustSortJSON(bz)
}

func (m *Faulters) MustUnmarshalBinaryBare(bz []byte) {
	if err := json.Unmarshal(bz, m); err != nil {
		panic(err) // handle the error according to your use case
	}
}
