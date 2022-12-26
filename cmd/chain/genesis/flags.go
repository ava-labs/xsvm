// Copyright (C) 2022, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package genesis

import (
	"math"
	"time"

	"github.com/spf13/pflag"

	"github.com/ava-labs/avalanchego/genesis"
	"github.com/ava-labs/avalanchego/ids"

	xsgenesis "github.com/ava-labs/xsvm/genesis"
)

const (
	TimeKey    = "time"
	AddressKey = "address"
	BalanceKey = "balance"
)

func AddFlags(flags *pflag.FlagSet) {
	flags.Int64(TimeKey, time.Now().Unix(), "Unix timestamp to include in the genesis")
	flags.String(AddressKey, genesis.EWOQKey.Address().String(), "Address to fund in the genesis")
	flags.Uint64(BalanceKey, math.MaxUint64, "Amount to provide the funded address in the genesis")
}

func ParseFlags(flags *pflag.FlagSet, args []string) (*xsgenesis.Genesis, error) {
	if err := flags.Parse(args); err != nil {
		return nil, err
	}

	if err := flags.Parse(args); err != nil {
		return nil, err
	}

	timestamp, err := flags.GetInt64(TimeKey)
	if err != nil {
		return nil, err
	}

	addrStr, err := flags.GetString(AddressKey)
	if err != nil {
		return nil, err
	}

	addr, err := ids.ShortFromString(addrStr)
	if err != nil {
		return nil, err
	}

	balance, err := flags.GetUint64(BalanceKey)
	if err != nil {
		return nil, err
	}

	return &xsgenesis.Genesis{
		Timestamp: timestamp,
		Allocations: []xsgenesis.Allocation{
			{
				Address: addr,
				Balance: balance,
			},
		},
	}, nil
}
