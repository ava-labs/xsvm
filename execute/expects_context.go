// Copyright (C) 2022, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package execute

import (
	"github.com/ava-labs/xsvm/block"
	"github.com/ava-labs/xsvm/tx"
)

var _ tx.Visitor = (*TxExpectsContext)(nil)

func ExpectsContext(blk *block.Stateless) (bool, error) {
	t := TxExpectsContext{}
	for _, tx := range blk.Txs {
		if err := tx.Unsigned.Visit(&t); err != nil {
			return false, err
		}
	}
	return t.Result, nil
}

type TxExpectsContext struct {
	Result bool
}

func (*TxExpectsContext) Transfer(*tx.Transfer) error {
	return nil
}

func (*TxExpectsContext) Export(*tx.Export) error {
	return nil
}

func (t *TxExpectsContext) Import(*tx.Import) error {
	t.Result = true
	return nil
}
